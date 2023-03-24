package controllers

import (
	"gower/app"
	"gower/app/http/requests"
	"gower/services"
)

type ExcpController struct {
	app.Controller
}

var Excp = new(ExcpController)

// BadRequest 异常请求
func (e *ExcpController) BadRequest(req *requests.ExcpRequest) services.Response {
	return res.Ok("excp/400", e.data(req))
}

// NotFound 404 未找到资源
func (e *ExcpController) NotFound(req *requests.ExcpRequest) services.Response {
	return res.Ok("excp/404", e.data(req))
}

func (e *ExcpController) data(req *requests.ExcpRequest) app.Data {
	return app.Data{
		"msg":    req.Msg,
		"detail": req.Detail,
	}
}
