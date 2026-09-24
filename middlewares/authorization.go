package middlewares

import (
	"errors"

	"food_delivery/common"
	"food_delivery/component/appctx"

	"github.com/gin-gonic/gin"
)

func RequireAuthor(appCtx appctx.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		hasFound := false

		for _, item := range allowRoles {
			if u.GetRole() == item {
				hasFound = true
				break
			}
		}

		if !hasFound {
			panic(common.ErrNotPermission(errors.New("permission denied")))
		}

		c.Next()
	}
}
