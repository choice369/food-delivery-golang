package jwt

import (
	"time"

	"food_delivery/component/token_provider"

	"github.com/dgrijalva/jwt-go"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload token_provider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data token_provider.TokenPayload, expiry int) (*token_provider.Token, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &myClaims{
		Payload: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	token, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}
	return &token_provider.Token{
		Token:     token,
		ExpiresAt: expiry,
		CreatedAt: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(token string) (*token_provider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(token, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, token_provider.ErrInvalidToken
	}

	if !res.Valid {
		return nil, token_provider.ErrInvalidToken
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, token_provider.ErrInvalidToken
	}

	return &claims.Payload, nil
}
