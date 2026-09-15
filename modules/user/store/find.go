package userstore

import (
	"context"

	"food_delivery/common"

	usermodel "food_delivery/modules/user/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindUser(ctx context.Context, cond map[string]interface{}, moreInfo ...string) (*usermodel.User, error) {
	db := s.db.Table(usermodel.User{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var user usermodel.User

	if err := db.Where(cond).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &user, nil
}
