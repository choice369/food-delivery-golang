package restaurant_repository

import (
	"context"

	"food_delivery/common"
	restaurantmodel "food_delivery/modules/restaurant/model"
)

type ListRestaurantStore interface {
	ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging, moreKey ...string) ([]restaurantmodel.Restaurant, error)
}

//type LikeRestaurantStore interface {
//	GetRestaurantLikes(ctx context.Context, ids []int) (map[int]int, error)
//}

type listRestaurantRepo struct {
	store ListRestaurantStore
	// likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}

func (biz *listRestaurantRepo) ListRestaurant(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, err
	}

	//ids := make([]int, len(result))
	//
	//for i := range ids {
	//	ids[i] = result[i].Id
	//}

	//likeMap, err := biz.likeStore.GetRestaurantLikes(ctx, ids)
	//if err != nil {
	//	log.Println(err)
	//	return result, nil
	//}
	//
	//for i, item := range result {
	//	result[i].LikedCount = likeMap[item.Id]
	//}

	return result, err
}
