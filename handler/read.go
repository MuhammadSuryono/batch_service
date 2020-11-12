package handler

import (
	"git.lifewood.dev/common-service/response"
	"git.lifewood.dev/services/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleList godoc
// @Summary Get User
// @Description Get paginated list of users
// @Tags Users
// @Param Authorization header string true "Bearer <token>"
// @Param payload body UserRequest true "Payload body"
// @Accept  json
// @Produce  json
// @Success 200 {object} UserResponse
// @Failure 200 {object} UserResponse
// @Router /v1/users [get]
func (*SkeletonHandler) HandleList(c *gin.Context) {
	//TODO: change your bind request list
	var r model.SkeletonRequest
	c.BindJSON(&r)
	//TODO: your logic list here

	//TODO: adjust response list
	c.JSON(http.StatusOK,
		model.SkeletonResponse{
			BaseResponse: response.BaseResponse{},
		})
}

// HandleSearch godoc
// @Summary Search User
// @Description Search name of user, return paginated list of users
// @Tags Users
// @Param Authorization header string true "Bearer <token>"
// @Param payload body UserRequest true "Payload body"
// @Accept  json
// @Produce  json
// @Success 200 {object} UserResponse
// @Failure 200 {object} UserResponse
// @Router /v1/users/search/{query} [get]
func (*SkeletonHandler) HandleSearch(c *gin.Context) {
	//TODO: change your bind request search
	var r model.SkeletonRequest
	c.BindJSON(&r)
	//TODO: your logic search here

	//TODO: adjust response search
	c.JSON(http.StatusOK,
		model.SkeletonResponse{
			BaseResponse: response.BaseResponse{},
		})
}
