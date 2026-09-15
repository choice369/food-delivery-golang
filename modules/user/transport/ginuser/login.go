package ginuser

import (
	"net/http"

	"food_delivery/common"
	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	"food_delivery/component/token_provider/jwt"
	userbiz "food_delivery/modules/user/biz"
	usermodel "food_delivery/modules/user/model"
	userstore "food_delivery/modules/user/store"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, 60*60*24*30)
		account, err := biz.Login(c.Request.Context(), &loginUserData)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessRes(account))
	}
}
