package handler

import (
	"fmt"
	db2 "git.lifewood.dev/common-service/db"
	"git.lifewood.dev/services/skeleton/excel"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (*BatchHanlder) HandleCheck(c *gin.Context) {
	var uriIdBatch model.UriIdBatch
	var batch model.Batch

	c.BindUri(&uriIdBatch)

	getBatch := db2.GormDb.Where("id = ? ", uriIdBatch.IdBatch).First(&batch)
	if getBatch.RowsAffected <= 0 {
		c.JSON(http.StatusNotFound, model.CommonResponse{
			IsSuccess: "1",
			Message:   "Data batch not found!",
			Data:      fmt.Sprintf("id batch %d", uriIdBatch.IdBatch),
		})
		return
	}

	// TODO: Check file exists or duplicate
	resultCheck := ""
	statusCheck := true
	//util := utilscustome.UtilsCustome()

	// TODO: Check rows
	excelCheck := excel.ExcelReadWriteEngine(batch.File)
	if excelCheck.CountRows() != batch.ImageQty {
		resultCheck += "Record Not Match \n"
		statusCheck = false
	}

	if !statusCheck {
		update := db2.GormDb.Model(model.Batch{}).
			Where("id = ?", uriIdBatch.IdBatch).
			Update(map[string]interface{}{
				"check_status": "Not Pass",
				"status":       "Checked",
				"remark_check": resultCheck})

		if update.RowsAffected <= 0 {
			c.JSON(http.StatusBadRequest, model.CommonResponse{
				IsSuccess: "1",
				Message:   "Error update data status",
				Data:      batch,
			})
			return
		}

		c.JSON(http.StatusOK, model.CommonResponse{
			IsSuccess: "0",
			Message:   "Success update batch on action check",
			Data:      batch,
		})
		return
	}

	update := db2.GormDb.Table("batches").
		Where("id = ?", uriIdBatch.IdBatch).
		Update(map[string]interface{}{
			"check_status": "Pass",
			"status":       "Checked"})

	if update.RowsAffected <= 0 {
		c.JSON(http.StatusBadRequest, model.CommonResponse{
			IsSuccess: "1",
			Message:   "Error update data status",
			Data:      batch,
		})
		return
	}

	c.JSON(http.StatusOK, model.CommonResponse{
		IsSuccess: "0",
		Message:   "Success update batch on action check",
		Data:      batch,
	})
	return
}
