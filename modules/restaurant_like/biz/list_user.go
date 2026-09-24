package restaurant_like_biz

import (
	"context"

	"food_delivery/common"
	restaurant_like_model "food_delivery/modules/restaurant_like/model"
)

type ListUserLikeRestaurantStore interface {
	GetUsersLikeRestaurant(ctx context.Context, cond map[string]interface{}, filter *restaurant_like_model.Filter, paging *common.Paging, moreKeys ...string) ([]common.SimpleUser, error)
}

type listUserLikeRestaurantBiz struct {
	store ListUserLikeRestaurantStore
}

func NewListUserLikeRestaurantBiz(store ListUserLikeRestaurantStore) *listUserLikeRestaurantBiz {
	return &listUserLikeRestaurantBiz{store: store}
}

func (biz *listUserLikeRestaurantBiz) ListUsers(ctx context.Context, filter *restaurant_like_model.Filter, paging *common.Paging) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersLikeRestaurant(ctx, nil, filter, paging)
	if err != nil {
		return nil, common.ErrCannotListEntity(err)
	}

	return users, nil
}
