package model

import "git.lifewood.dev/common-service/response"

type SkeletonResponse struct {
	response.BaseResponse
	UserId int64 `json:"user_id"`
}
