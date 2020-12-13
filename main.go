package main

import (
	"git.lifewood.dev/common-service/db"
	_ "git.lifewood.dev/services/skeleton/docs"
	"git.lifewood.dev/services/skeleton/handler"
	"git.lifewood.dev/services/skeleton/middleware"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
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
	db.GormDb.AutoMigrate(&model.Batch{})
	db.GormDb.AutoMigrate(&model.SmallBatch{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, model.CommonResponse{
			IsSuccess: "1",
			Message:   context.Request.Host + context.Request.RequestURI + " Route Not Found",
			Data:      nil,
		})
	})

	r.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusMethodNotAllowed, model.CommonResponse{
			IsSuccess: "1",
			Message:   "Metode Not Allowed!",
			Data:      nil,
		})
	})

	r.MaxMultipartMemory = 8 << 20
	//r.Static("/", "./upload")

	handle := handler.NewBatchHanlder()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", handle.HandlePing)
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// TODO: update your handler
		v1.Use(middleware.AuthorizeHeader())
		{
			v1.GET("/batchs", handle.HandleAllBatch)
			v1.GET("/batchs/project/:projectCode", handle.HandleProjectBatch)
			v1.GET("/batchs/check/:idBatch", handle.HandleCheck)
			v1.POST("/batchs", handle.HandleCreate)
		}
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		r.Run()
	}

	r.Run(":" + port)

}
