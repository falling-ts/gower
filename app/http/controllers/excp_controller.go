package controllers

import (
	"gitee.com/falling-ts/gower/app"
	"gitee.com/falling-ts/gower/app/http/requests"
	"gitee.com/falling-ts/gower/services"
)

type ExcpController struct {
	app.Controller
}

var Excp = new(ExcpController)

// BadRequest 异常请求
func (e *ExcpController) BadRequest(req *requests.ExcpRequest) services.Response {
	return res.Ok("exc/400", e.data(req))
}

// NotFound 404 未找到资源
func (e *ExcpController) NotFound(req *requests.ExcpRequest) services.Response {
	return res.Ok("exc/404", e.data(req))
}

func (e *ExcpController) data(req *requests.ExcpRequest) app.Data {
	return app.Data{
		"msg":    req.Msg,
		"detail": req.Detail,
	}
}
