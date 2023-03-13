package controllers

import (
	"gower/app"
	"gower/services"
)

type Controllers struct{}

// HTML Data
type data map[string]any

// ResponseData 非 HTML 请求的成功响应体
type ResponseData struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var routeSrv = app.Route()

// 通用响应方法
func (c Controllers) response(args ...any) services.Response {
	return routeSrv.Response(new(ResponseData), args...)
}

// Set 设置响应体内容
func (r *ResponseData) Set(arg any) services.ResponseData {
	switch arg.(type) {
	case int:
		r.Code = arg.(int)
	case string:
		r.Msg = arg.(string)
	default:
		r.Data = arg
	}
	return r
}
