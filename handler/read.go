package handler

import (
	db2 "git.lifewood.dev/common-service/db"
	"git.lifewood.dev/services/skeleton/model"
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
func (*BatchHanlder) HandleAllBatch(c *gin.Context) {
	//TODO: change your bind request list
	var queryRequest model.QueryRequest
	c.BindQuery(&queryRequest)

	//TODO: your logic list here
	db := db2.GormDb
	dataBatch := model.PaginationBatch(db, queryRequest)

	//TODO: adjust response list
	c.JSON(http.StatusOK, dataBatch)
}

func (*BatchHanlder) HandleProjectBatch(c *gin.Context) {
	//TODO: change your bind request list
	var queryRequest model.QueryRequest
	var uriCodeProject model.UriProjectId
	c.BindQuery(&queryRequest)
	c.BindUri(&uriCodeProject)

	//TODO: your logic list here
	db := db2.GormDb.Where("id_project = ? ", uriCodeProject.ProjectCode)
	dataBatch := model.PaginationBatch(db, queryRequest)

	//TODO: adjust response list
	c.JSON(http.StatusOK, dataBatch)
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
func (*BatchHanlder) HandleSearch(c *gin.Context) {
	//TODO: change your bind request search
	//var r model.SkeletonRequest
	//c.BindJSON(&r)
	////TODO: your logic search here
	//
	////TODO: adjust response search
	//c.JSON(http.StatusOK,
	//	model.SkeletonResponse{
	//		BaseResponse: response.BaseResponse{},
	//	})
}
