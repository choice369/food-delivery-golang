package appctx

import (
	"food_delivery/component/upload_provider"
	"food_delivery/pubsub"
	"food_delivery/skio"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	UploadProvider() upload_provider.UploadProvider
	SecretKey() string
	GetPubSub() pubsub.Pubsub
}

type appCtx struct {
	db             *gorm.DB
	uploadProvider upload_provider.UploadProvider
	secretKey      string
	ps             pubsub.Pubsub
	rtEngine       skio.RealtimeEngine
}

func NewAppContext(db *gorm.DB, upload_provider upload_provider.UploadProvider, secretKey string, ps pubsub.Pubsub) *appCtx {
	return &appCtx{db: db, uploadProvider: upload_provider, secretKey: secretKey, ps: ps}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) UploadProvider() upload_provider.UploadProvider {
	return ctx.uploadProvider
}

func (ctx *appCtx) SecretKey() string { return ctx.secretKey }

func (ctx *appCtx) GetPubSub() pubsub.Pubsub {
	return ctx.ps
}

func (ctx *appCtx) SetRealtimeEngine(rtEngine skio.RealtimeEngine) { ctx.rtEngine = rtEngine }
