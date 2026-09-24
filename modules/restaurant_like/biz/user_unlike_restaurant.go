package restaurant_like_biz

import (
	"context"

	restaurant_like_model "food_delivery/modules/restaurant_like/model"
)

type UserUnlikeRestaurantStore interface {
	Delete(ctx context.Context, userId int, restaurantId int) error
}

type DecLikedCountResStore interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnlikeRestaurantBiz struct {
	store    UserUnlikeRestaurantStore
	decStore DecLikedCountResStore
}

func NewUserUnlikeRestaurantBiz(store UserUnlikeRestaurantStore, decStore DecLikedCountResStore) *userUnlikeRestaurantBiz {
	return &userUnlikeRestaurantBiz{store, decStore}
}

func (biz *userUnlikeRestaurantBiz) DislikeRestaurant(ctx context.Context, userId int, restaurantId int) error {
	err := biz.store.Delete(ctx, userId, restaurantId)
	if err != nil {
		return restaurant_like_model.ErrCannotUnLikeRestaurant(err)
	}

	//j := async_job.NewJob(func(ctx context.Context) error {
	//	return biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	//})
	//
	//async_job.NewGroup(true, j).Run(ctx)

	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.decStore.DecreaseLikeCount(ctx, restaurantId); err != nil {
	//		log.Println(err)
	//
	//		for i := 1; i <= 5; i++ {
	//			err := biz.decStore.DecreaseLikeCount(ctx, restaurantId)
	//			if err == nil {
	//				break
	//			}
	//			time.Sleep(3 * time.Second)
	//		}
	//	}
	//}()

	return nil
}
