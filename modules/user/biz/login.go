package userbiz

import (
	"context"

	"food_delivery/component/token_provider"

	"food_delivery/common"

	"food_delivery/component/appctx"

	usermodel "food_delivery/modules/user/model"
)

type LoginStorage interface {
	FindUser(ctx context.Context, cond map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx        appctx.AppContext
	storeUser     LoginStorage
	tokenProvider token_provider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storeUser LoginStorage, tokenProvider token_provider.Provider, hasher Hasher, expiry int) *loginBusiness {
	return &loginBusiness{
		storeUser:     storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*token_provider.Token, error) {
	user, err := business.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := business.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	payload := token_provider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := business.tokenProvider.Generate(payload, business.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
