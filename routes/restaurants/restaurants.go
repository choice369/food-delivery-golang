package restaurants

import (
	"food_delivery/component/appctx"
	"food_delivery/middlewares"
	"food_delivery/modules/restaurant/transport/ginrestaurant"
	"food_delivery/modules/restaurant_like/transport/ginrstlike"

	"github.com/gin-gonic/gin"
)

func SetupRestaurantsRoutes(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	restaurants := v1.Group("/restaurants", middlewares.RequireAuthen(appCtx))

	restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))

	restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

	restaurants.POST("/:id/like", ginrstlike.UserLikeRestaurant(appCtx))

	restaurants.DELETE("/:id/dislike", ginrstlike.UserUnlikeRestaurant(appCtx))

	//v1.GET("/restaurants/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//	var restaurant Restaurant
	//
	//	db.Where("id = ?", id).First(&restaurant)
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": restaurant,
	//	})
	//})

	//v1.PATCH("/restaurants/:id", func(c *gin.Context) {
	//	var data RestaurantUpdate
	//	if err := c.ShouldBind(&data); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	db.Where("id = ?", data.OwnerId).Updates(&data)
	//})
}
