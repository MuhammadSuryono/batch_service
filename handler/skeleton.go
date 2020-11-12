package handler

import "github.com/gin-gonic/gin"

//TODO: Adjust your Handler
type ISkeletonHandler interface {
	HandleCreate(c *gin.Context)
}
type SkeletonHandler struct {
}

func NewSkeletonHandler() ISkeletonHandler {
	return &SkeletonHandler{}
}
