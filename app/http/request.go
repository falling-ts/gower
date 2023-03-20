package http

import "github.com/gin-gonic/gin"

// Request 通用请求接口
type Request interface {
	Validate(ctx *gin.Context, req Request) error
	SetContext(c *gin.Context)
}
