package main

import (
	"git.lifewood.dev/common-service/db"
	_ "git.lifewood.dev/services/skeleton/docs"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/joho/godotenv"
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title LiFT - User Service
// @version 0.0.1
// @termsOfService http://swagger.io/terms/
// @description Handler service for user

// @contact.name Husnan
// @contact.url https://lifewood.com
// @contact.email emhusnan@lifewood.com

// @license.name LiFT Lifewood
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:7000
// @BasePath /

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization "Bearer: xxxxxx"
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
	db.GormDb.AutoMigrate(&model.SkeletonRequest{})
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	v1 := r.Group("/v1")
	{
		// TODO: update your handler
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, "ok")
		})
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":7000")

}
