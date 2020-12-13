package model

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/jinzhu/gorm"
)

type CommonResponse struct {
	IsSuccess string      `json:"is_success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func PaginationBatch(db *gorm.DB, query QueryRequest) *pagination.Paginator {
	var batch []Batch
	paginator := pagination.Paging(&pagination.Param{
		DB:      db,
		Page:    query.Page,
		Limit:   query.PerPage,
		OrderBy: []string{"id desc"},
		ShowSQL: true,
	}, &batch)

	return paginator
}
