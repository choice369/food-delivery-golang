package ginrestaurant

import (
	"net/http"

	restaurant_repository "food_delivery/modules/restaurant/repository"

	restaurantbiz "food_delivery/modules/restaurant/biz"
	restaurantstorage "food_delivery/modules/restaurant/storage"

	restaurantmodel "food_delivery/modules/restaurant/model"

	"food_delivery/common"
	"food_delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func ListRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fullfill()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var result []restaurantmodel.Restaurant

		store := restaurantstorage.NewSqlStore(db)
		//likeStore := restaurant_like_storage.NewSqlStore(db)
		repo := restaurant_repository.NewListRestaurantRepo(store)
		biz := restaurantbiz.NewListRestaurantBiz(repo)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessRes(result, pagingData, filter))
	}
}
