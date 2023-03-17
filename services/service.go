package services

import "github.com/gin-gonic/gin"

// Service 服务通用接口
type Service interface {
	Init(...any)
}

// Accepts 接受的响应数据类型
var Accepts = []string{
	gin.MIMEJSON,
	gin.MIMEHTML,
	gin.MIMEXML,
	gin.MIMEYAML,
	gin.MIMETOML,
	gin.MIMEPlain,
}
