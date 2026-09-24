package userroutes

import (
	"food_delivery/component/appctx"
	"food_delivery/middlewares"
	"food_delivery/modules/upload/transport/ginupload"
	"food_delivery/modules/user/transport/ginuser"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(appCtx appctx.AppContext, v1 *gin.RouterGroup) {
	u := v1.Group("/u")

	u.POST("upload", ginupload.UploadImage(appCtx))

	u.POST("register", ginuser.Register(appCtx))

	u.POST("authentication", ginuser.Login(appCtx))

	u.GET("/profile", middlewares.RequireAuthen(appCtx), ginuser.Profile(appCtx))
}
