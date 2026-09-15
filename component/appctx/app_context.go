package appctx

import (
	"food_delivery/component/upload_provider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() upload_provider.UploadProvider
	SecretKey() string
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider upload_provider.UploadProvider
	secretKey      string
}

func NewAppContext(db *gorm.DB, upload_provider upload_provider.UploadProvider, secretKey string) *appCtx {
	return &appCtx{db: db, uploadProvider: upload_provider, secretKey: secretKey}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() upload_provider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }
