package handler

import (
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func (*BatchHanlder) HandlePing(context *gin.Context) {
	context.JSON(http.StatusOK, model.CommonResponse{
		IsSuccess: "0",
		Message:   "Success ping",
		Data:      os.Getenv("APP_NAME"),
	})
}
