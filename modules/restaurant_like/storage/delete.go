package restaurant_like_storage

import (
	"context"

	"food_delivery/common"
	restaurant_like_model "food_delivery/modules/restaurant_like/model"
)

func (s *sqlStore) Delete(ctx context.Context, restaurantId int, userId int) error {
	db := s.db

	if err := db.Table(restaurant_like_model.Like{}.TableName()).
		Where("user_id = ? and restaurant_id = ?", userId, restaurantId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
