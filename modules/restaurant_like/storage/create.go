package restaurant_like_storage

import (
	"context"

	"food_delivery/common"

	restaurant_like_model "food_delivery/modules/restaurant_like/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurant_like_model.Like) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
