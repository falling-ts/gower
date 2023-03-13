package controllers

import (
	"gower/app"
	"gower/app/responses"
	"gower/services"
)

type Controllers struct{}

// HTML Data
type data map[string]any

var routeSrv = app.Route()

// 通用响应方法
func (c Controllers) response(args ...any) services.Response {
	return routeSrv.Response(new(responses.Responses), args...)
}
