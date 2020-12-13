package handler

import (
	"fmt"
	"git.lifewood.dev/common-service/db"
	"git.lifewood.dev/services/skeleton/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func (*BatchHanlder) ReallocateBatch(c *gin.Context) {
	var requestReallocate model.RequestReallocate

	err := c.BindJSON(&requestReallocate)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.CommonResponse{
			IsSuccess: "1",
			Message:   fmt.Sprintf("Error on body payload %s", err.Error()),
			Data:      requestReallocate,
		})
		return
	}

	var isSuccess = true
	err = db.GormDb.Transaction(func(tx *gorm.DB) error {
		for _, id := range requestReallocate.IdBatch {
			if err := tx.Model(&model.Batch{}).Where("id = ? ", id).Update(map[string]interface{}{"id_site": requestReallocate.IdSite}).Error; err != nil {
				isSuccess = false
				return err
			}
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, model.CommonResponse{
			IsSuccess: "1",
			Message:   fmt.Sprintf("Error when reallocate data batch %s", err.Error()),
			Data:      requestReallocate,
		})
		return
	}

	c.JSON(http.StatusOK, model.CommonResponse{
		IsSuccess: "0",
		Message:   "Success reallocate data batch",
		Data:      requestReallocate,
	})
}
