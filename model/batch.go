package model

import (
	"github.com/jinzhu/gorm"
)

type Batch struct {
	IdProject   int64  `gorm:"type:BIGINT" json:"id_project"`
	IdSite int64 `gorm:"type:BIGINT" json:"id_site"`
	BatchName   string `gorm:"type:VARCHAR(20)" json:"batch_name"`
	ImageQty    int64  `gorm:"type:BIGINT" json:"image_qty"`
	Split       int64  `gorm:"type:BIGINT" json:"split"`
	Priority    int    `gorm:"type:INT" json:"priority"`
	Status      string `gorm:"type:VARCHAR(25)" json:"status"`
	CreatedBy   int64  `gorm:"type:BIGINT" json:"created_by"`
	PostedBy    int64  `gorm:"type:BIGINT" json:"posted_by"`
	CheckDate   string `json:"check_date"`
	PostedDate  string `json:"posted_date"`
	RemarkCheck string `gorm:"type:TEXT" json:"remark_check"`
	CheckStatus string `gorm:"type:VARCHAR(20)" json:"check_status"`
	File        string `gorm:"type:VARCHAR(100)" json:"file"`
	gorm.Model
}
