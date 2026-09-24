package ginrstlike

import (
	"net/http"

	"food_delivery/common"
	"food_delivery/component/appctx"
	restaurant_like_biz "food_delivery/modules/restaurant_like/biz"
	restaurant_like_model "food_delivery/modules/restaurant_like/model"
	restaurant_like_storage "food_delivery/modules/restaurant_like/storage"

	"github.com/gin-gonic/gin"
)

func ListUsers(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter := restaurant_like_model.Filter{
			RestaurantId: int(uid.LocalID()),
		}

		var paging common.Paging

		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fullfill()

		store := restaurant_like_storage.NewSqlStore(appCtx.GetMainDBConnection())
		biz := restaurant_like_biz.NewListUserLikeRestaurantBiz(store)

		result, err := biz.ListUsers(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessRes(result, paging, filter))
	}
}
