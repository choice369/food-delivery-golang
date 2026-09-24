package ginrstlike

import (
	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurant_like_biz "food_delivery/modules/restaurant_like/biz"
	restaurant_like_model "food_delivery/modules/restaurant_like/model"
	restaurant_like_storage "food_delivery/modules/restaurant_like/storage"

	"github.com/gin-gonic/gin"
)

// POST /v1/restaurant/:id/like
// POST /v1/restaurant-likes (not recommend)

func UserLikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := restaurant_like_model.Like{
			RestaurantId: int(uid.LocalID()),
			UserId:       requester.GetUserId(),
		}

		store := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
		//incStore := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurant_like_biz.NewUserLikeRestaurantBiz(store, appCtx.GetPubSub())

		if err := biz.LikeRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessRes(true))
	}
}
