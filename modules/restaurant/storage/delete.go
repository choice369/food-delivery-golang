package restaurantstorage

import (
	"context"

	restaurantmodel "food_delivery/modules/restaurant/model"
)

func (s *sqlStore) Delete(ctx context.Context, id uint32) error {
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
