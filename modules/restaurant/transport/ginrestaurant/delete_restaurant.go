package ginrestaurant

import (
	"net/http"

	"food_delivery/common"

	"food_delivery/component/appctx"

	restaurantbiz "food_delivery/modules/restaurant/biz"
	restaurantstorage "food_delivery/modules/restaurant/storage"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		//id, err := strconv.Atoi(c.Param("id"))

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store, requester)

		if err := biz.DeleteRestaurant(c.Request.Context(), uid.LocalID()); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(true))
	}
}
