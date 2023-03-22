package services

import "github.com/gin-gonic/gin"

// Exception 异常内容接口
type Exception interface {
	error
	Service
	Set(arg any) Exception
	Get(field string) (any, error)
	New(code int, args ...any) Exception
	Build(args ...any) Exception
	Handle(c *gin.Context) bool
}

// ExceptionService 异常服务接口
type ExceptionService interface {
	Service
	Build(args ...any) Exception
	Handle(c *gin.Context) bool
}
