package handler

import (
	"git.lifewood.dev/common-service/response"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleCreate godoc
// @Summary Create new user
// @Description Create new user
// @Tags Users
// @Param Authorization header string true "Bearer <token>"
// @Param payload body  model.UserRequest true "Payload body"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserResponse
// @Failure 200 {object} model.UserResponse
// @Router /v1/users [post]
func (*SkeletonHandler) HandleCreate(c *gin.Context) {
	//TODO: change your bind request create
	var r model.SkeletonRequest
	c.BindJSON(&r)
	//TODO: your logic create here

	//TODO: adjust response create
	c.JSON(http.StatusOK,
		model.SkeletonResponse{
			BaseResponse: response.BaseResponse{},
		})

}

