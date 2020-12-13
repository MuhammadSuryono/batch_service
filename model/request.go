package model

import (
	"encoding/base64"
	"encoding/json"
)

type QueryRequest struct {
	Page    int `form:"page"`
	PerPage int `form:"per_page"`
}

type PayloadRequest struct {
	Payload string `form:"payload"`
}

type UriProjectId struct {
	ProjectCode int64 `uri:"projectCode"`
}

type RequestCreate struct {
	IdProject int64  `json:"id_project" binding:"required"`
	IdSite    int64  `json:"id_site" binding:"required"`
	ImageQty  int64  `json:"image_qty" binding:"required"`
	BatchName string `json:"batch_name" binding:"required"`
	Split     int64  `json:"split" binding:"required"`
	Priority  int    `json:"priority" binding:"required"`
	UserId    int64  `json:"user_id" binding:"required"`
}

type UriIdBatch struct {
	IdBatch int64 `uri:"idBatch"`
}

type RequestReallocate struct {
	IdBatch []int64 `json:"id_batch" binding:"required"`
	IdSite  int64   `json:"id_site" binding:"required"`
}

func EncodeJsonMarshal(data interface{}) string {
	jsonReq, _ := json.Marshal(data)
	enc := base64.StdEncoding.EncodeToString([]byte(string(jsonReq)))
	return enc
}

func DecodePayloadQuery(b64 string, model interface{}) {
	sDec, _ := base64.StdEncoding.DecodeString(b64)
	json.Unmarshal(sDec, &model)
}
