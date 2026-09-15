package restaurantstorage

import (
	"context"

	"food_delivery/common"

	restaurantmodel "food_delivery/modules/restaurant/model"
)

// Step 1: Fetch 1000 ids => caching
// Step 2: Fetch 50 first ids => data (page 1)
// Step 3: Page 2, omit 50 ids and fetch the next 50 ids

func (s *sqlStore) ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKey ...string) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id=?", f.OwnerId)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return result, err
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			return result, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.LocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := s.db.Limit(paging.Limit).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
