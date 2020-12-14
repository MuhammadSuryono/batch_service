package handler

import (
	"fmt"
	"git.lifewood.dev/common-service/db"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
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
func (*BatchHanlder) HandleCreate(c *gin.Context) {
	//TODO: change your bind request create
	var payload model.PayloadRequest
	var requestCreate model.RequestCreate

	c.BindQuery(&payload)
	model.DecodePayloadQuery(payload.Payload, &requestCreate)

	////TODO: your logic create here
	file, _ := c.FormFile("file")
	filename := filepath.Base(file.Filename)

	err := c.SaveUploadedFile(file, filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.CommonResponse{
			IsSuccess: "1",
			Message:   fmt.Sprintf("upload file err: %s", err.Error()),
			Data:      filename,
		})
		return
	}

	checkBatchName := db.GormDb.Where("batch_name = ? AND check_status = ?", requestCreate.BatchName, "Pass").First(&model.Batch{})
	if checkBatchName.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, model.CommonResponse{
			IsSuccess: "1",
			Message:   fmt.Sprintf("Batch name %s is exists", requestCreate.BatchName),
			Data:      requestCreate.BatchName,
		})
		return
	}

	batch := &model.Batch{
		IdProject: requestCreate.IdProject,
		IdSite: requestCreate.IdSite,
		BatchName: requestCreate.BatchName,
		ImageQty:  requestCreate.ImageQty,
		Split:     requestCreate.Split,
		Priority:  requestCreate.Priority,
		Status:    "Uploaded",
		CreatedBy: requestCreate.UserId,
		File:      filename,
	}

	create := db.GormDb.Create(&batch)
	if create.RowsAffected <= 0 {
		c.JSON(http.StatusBadRequest, model.CommonResponse{
			IsSuccess: "1",
			Message:   "Cannot save data batch, error on " + create.Error.Error(),
			Data:      requestCreate,
		})
		return
	}

	////TODO: adjust response create
	c.JSON(http.StatusBadRequest, model.CommonResponse{
		IsSuccess: "0",
		Message:   "Save data batch success ",
		Data:      requestCreate,
	})
	return
}
