package ginuser

import (
	"context"
	"net/http"

	"food_delivery/common"

	"food_delivery/component/appctx"
	"food_delivery/component/hasher"
	userbiz "food_delivery/modules/user/biz"
	usermodel "food_delivery/modules/user/model"
	userstore "food_delivery/modules/user/store"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()

		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(context.Background(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessRes(data.FakeId.String()))
	}
}
