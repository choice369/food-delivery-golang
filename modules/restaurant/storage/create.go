package restaurantstorage

import (
	"context"

	"food_delivery/common"

	restaurantmodel "food_delivery/modules/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
