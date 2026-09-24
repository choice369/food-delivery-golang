package restaurant_like_biz

import (
	"context"
	"log"

	"food_delivery/common"

	"food_delivery/pubsub"

	restaurant_like_model "food_delivery/modules/restaurant_like/model"
)

type UserLikeRestaurantStore interface {
	Create(ctx context.Context, data *restaurant_like_model.Like) error
}

//type IncLikedCountResStore interface {
//	IncreaseLikeCount(ctx context.Context, id int) error
//}

type userLikeRestaurantBiz struct {
	store UserLikeRestaurantStore
	// incStore IncLikedCountResStore
	ps pubsub.Pubsub
}

func NewUserLikeRestaurantBiz(store UserLikeRestaurantStore, ps pubsub.Pubsub) *userLikeRestaurantBiz {
	return &userLikeRestaurantBiz{store: store, ps: ps}
}

func (biz *userLikeRestaurantBiz) LikeRestaurant(ctx context.Context, data *restaurant_like_model.Like) error {
	err := biz.store.Create(ctx, data)
	if err != nil {
		return restaurant_like_model.ErrCannotLikeRestaurant(err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUserLikeRestaurant, pubsub.NewMessage(data)); err != nil {
		log.Println(err)
	}
	//
	//j := async_job.NewJob(func(ctx context.Context) error {
	//	return biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId)
	//})
	//
	//async_job.NewGroup(true, j).Run(ctx)

	//go func() {
	//	defer common.AppRecover()
	//	if err := biz.incStore.IncreaseLikeCount(ctx, data.RestaurantId); err != nil {
	//		log.Println(err)
	//	}
	//}()

	return nil
}
