package controllers

import (
	"gower/app"
	"gower/services"
)

type Controller struct{}

// Data HTML 数据
type Data map[string]any

var res = app.Response()

// 200 ok
func (c Controller) ok(args ...any) services.Response {
	return res.Ok(args...)
}
