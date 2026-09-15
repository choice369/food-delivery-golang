package restaurantbiz

import (
	"context"
	"errors"

	"food_delivery/common"

	restaurantmodel "food_delivery/modules/restaurant/model"
)

type DeleteRestaurantStore interface {
	Delete(ctx context.Context, id uint32) error
	FindDataWithCondition(ctx context.Context, condition map[string]interface{}, moreKey ...string) (*restaurantmodel.Restaurant, error)
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id uint32) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(restaurantmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil)
	}

	if oldData.Status == 0 {
		return errors.New("restaurant has been deleted")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
