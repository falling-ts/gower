package services

import "github.com/gin-gonic/gin"

// Response 响应体接口
type Response interface {
	Service
	Set(any) Response
	New(code int, args ...any) Response
	Build(code int, args ...any) Response
	Handle(c *gin.Context) bool
}

// ResponseService 响应体服务接口
type ResponseService interface {
	Service
	Build(code int, args ...any) Response
	Handle(c *gin.Context) bool
}
