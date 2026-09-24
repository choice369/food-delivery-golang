package subscriber

import (
	"context"

	"food_delivery/common"
	restaurant_like_storage "food_delivery/modules/restaurant_like/storage"

	"food_delivery/component/appctx"
)

func DecreaseLikeCountAfterUserDisLikeRestaurant(appCtx appctx.AppContext, ctx context.Context) {
	c, _ := appCtx.GetPubSub().Subscribe(ctx, common.TopicUserLikeRestaurant)

	store := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())

	go func() {
		defer common.AppRecover()
		for {
			msg := <-c
			likeData := msg.Data().(HasRestaurantId)
			_ = store.DecreaseLikeCount(ctx, likeData.GetRestaurantId())
		}
	}()
}
