package adminroutes

import (
	"food_delivery/component/appctx"
	"food_delivery/middlewares"
	"food_delivery/modules/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin", middlewares.RequireAuthen(appCtx), middlewares.RequireAuthor(appCtx, "admin", "mod"))

	{
		admin.GET("/profile", ginuser.Profile(appCtx))
	}
}
