package restaurantbiz

import (
	"context"

	"food_delivery/common"
	restaurantmodel "food_delivery/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKey ...string) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, err
}
