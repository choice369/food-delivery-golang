package main

import (
	"log"
	"os"

	"food_delivery/skio"

	"food_delivery/subscriber"

	"food_delivery/pubsub/local_pubsub"

	adminroutes "food_delivery/routes/admin"
	restaurantsroutes "food_delivery/routes/restaurants"
	userroutes "food_delivery/routes/user"

	"food_delivery/component/upload_provider"

	"food_delivery/middlewares"

	"food_delivery/component/appctx"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()

	dsn := os.Getenv("MYSQL_CONF_STRING")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_API_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")
	secretKey := os.Getenv("SECRET_KEY")

	s3Provider := upload_provider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = db.Debug()

	ps := local_pubsub.NewLocalPubSub()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)

	r := gin.Default()

	r.Static("/static", "./static")

	appContext := appctx.NewAppContext(db, s3Provider, secretKey, ps)

	//subscriber.Setup(appContext, context.Background())
	subscriber.NewEngine(appContext).Start()

	v1 := r.Group("/v1", middlewares.Recover(appContext))

	restaurantsroutes.SetupRestaurantsRoutes(appContext, v1)

	userroutes.SetupUserRoutes(appContext, v1)

	adminroutes.SetupAdminRoutes(appContext, v1)

	rtEngine := skio.NewEngine()
	appContext.SetRealtimeEngine(rtEngine)

	rtEngine.Run(appContext, r)

	r.Run()
}
