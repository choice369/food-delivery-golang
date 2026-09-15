package main

import (
	"log"
	"os"

	"food_delivery/modules/user/transport/ginuser"

	"food_delivery/component/upload_provider"

	"food_delivery/modules/upload/transport/ginupload"

	"food_delivery/middlewares"

	"food_delivery/component/appctx"

	"food_delivery/modules/restaurant/transport/ginrestaurant"

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

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)

	r := gin.Default()

	r.Static("/static", "./static")

	appContext := appctx.NewAppContext(db, s3Provider, secretKey)

	v1 := r.Group("/v1", middlewares.Recover(appContext))

	v1.POST("/restaurant", ginrestaurant.CreateRestaurant(appContext))

	v1.POST("upload", ginupload.UploadImage(appContext))

	v1.POST("register", ginuser.Register(appContext))

	v1.POST("authentication", ginuser.Login(appContext))

	v1.GET("/profile", middlewares.RequireAuth(appContext), ginuser.Profile(appContext))

	//v1.GET("/restaurants/:id", func(c *gin.Context) {
	//	id, err := strconv.Atoi(c.Param("id"))
	//	if err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//	}
	//	var restaurant Restaurant
	//
	//	db.Where("id = ?", id).First(&restaurant)
	//	c.JSON(http.StatusOK, gin.H{
	//		"data": restaurant,
	//	})
	//})

	v1.GET("/restaurants", ginrestaurant.ListRestaurant(appContext))

	//v1.PATCH("/restaurants/:id", func(c *gin.Context) {
	//	var data RestaurantUpdate
	//	if err := c.ShouldBind(&data); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{
	//			"error": err.Error(),
	//		})
	//		return
	//	}
	//
	//	db.Where("id = ?", data.OwnerId).Updates(&data)
	//})

	v1.DELETE("/restaurants/:id", ginrestaurant.DeleteRestaurant(appContext))

	r.Run()
}
