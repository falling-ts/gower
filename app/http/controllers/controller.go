package controllers

import (
	"gower/app"
	"gower/app/responses"
	"gower/services"
)

type Controller struct{}

// Data HTML 数据
type Data map[string]any

var routeSrv = app.Route()

// 通用响应方法
func (c Controller) response(args ...any) services.Response {
	return routeSrv.Response(new(responses.Responses), args...)
}
