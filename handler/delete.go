package handler

import (
	"git.lifewood.dev/common-service/response"
	"git.lifewood.dev/services/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleDelete godoc
// @Summary Delete User
// @Description Delete spesific user
// @Tags Users
// @Param Authorization header string true "Bearer <token>"
// @Param payload body  UserRequest true "Payload body"
// @Accept  json
// @Produce  json
// @Success 200 {object} UserResponse
// @Failure 200 {object} UserResponse
// @Router /v1/users/{id} [delete]
func (*SkeletonHandler) HandleDelete(c *gin.Context) {
	//TODO: change your bind request delete
	var r model.SkeletonRequest
	c.BindJSON(&r)
	//TODO: your logic delete here

	//TODO: adjust response delete
	c.JSON(http.StatusOK,
		model.SkeletonResponse{
			BaseResponse: response.BaseResponse{},
		})
}
