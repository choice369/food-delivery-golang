package ginrstlike

import (
	"net/http"

	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurant_like_biz "food_delivery/modules/restaurant_like/biz"
	restaurant_like_storage "food_delivery/modules/restaurant_like/storage"

	"github.com/gin-gonic/gin"
)

func UserUnlikeRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		store := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
		decStore := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurant_like_biz.NewUserUnlikeRestaurantBiz(store, decStore)

		if err := biz.DislikeRestaurant(c.Request.Context(), requester.GetUserId(), int(uid.LocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
