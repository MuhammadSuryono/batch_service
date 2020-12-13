package handler

import "github.com/gin-gonic/gin"

//TODO: Adjust your Handler
type IBatchHanlder interface {
	HandlePing(c *gin.Context)
	HandleCreate(c *gin.Context)
	HandleAllBatch(c *gin.Context)
	HandleProjectBatch(c *gin.Context)
	HandleCheck(c *gin.Context)
}

type BatchHanlder struct {
}

func NewBatchHanlder() IBatchHanlder {
	return &BatchHanlder{}
}
