package subscriber

import (
	"context"

	"food_delivery/pubsub"

	restaurant_like_storage "food_delivery/modules/restaurant_like/storage"

	"food_delivery/component/appctx"
)

type HasRestaurantId interface {
	GetRestaurantId() int
	// GetUserId() int
}

func IncreaseLikeCountAfterUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "IncreaseLikeCountAfterUserLikeRestaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}

func PushNotificationWhenUserLikeRestaurant(appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "IncreaseLikeCountAfterUserLikeRestaurant",
		Hdl: func(ctx context.Context, message *pubsub.Message) error {
			store := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasRestaurantId)
			return store.IncreaseLikeCount(ctx, likeData.GetRestaurantId())
		},
	}
}
