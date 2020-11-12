package handler

import (
	"git.lifewood.dev/common-service/response"
	"git.lifewood.dev/services/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleUpdate godoc
// @Summary Update User
// @Description Update spesific user
// @Tags Users
// @Param Authorization header string true "Bearer <token>"
// @Param payload body  UserRequest true "Payload body"
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserResponse
// @Failure 200 {object} model.UserResponse
// @Router /v1/users/{id} [put]
func (*SkeletonHandler) HandleUpdate(c *gin.Context) {
	//TODO: change your bind request update
	var r model.SkeletonRequest
	c.BindJSON(&r)
	//TODO: your logic update here

	//TODO: adjust update response
	c.JSON(http.StatusOK,
		model.SkeletonResponse{
			BaseResponse: response.BaseResponse{},
		})
}
