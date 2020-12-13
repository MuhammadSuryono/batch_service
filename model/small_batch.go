package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SmallBatch struct {
	IdBatch    int64     `gorm:"type:BIGINT" json:"id_batch"`
	ImageQty   int64     `gorm:"type:BIGINT" json:"image_qty"`
	Status     string    `gorm:"type:VARCHAR(25)" json:"status"`
	Priority   int       `json:"priority" form:"priority"`
	StatusDate time.Time `json:"status_date"`
	gorm.Model
}
