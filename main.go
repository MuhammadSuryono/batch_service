package main

import (
	"git.lifewood.dev/common-service/db"
	_ "git.lifewood.dev/services/user/docs"
	"git.lifewood.dev/services/user/model"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"
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
	db.GormDb.AutoMigrate(&model.User{})
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	r.Group("/v1", func(context *gin.Context) {
		// TODO: update your handler
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":7000")

}
