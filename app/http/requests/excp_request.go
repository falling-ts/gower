package requests

import "gower/app"

type ExcpRequest struct {
	app.Request
	Msg    string `form:"msg" binding:"required"`
	Detail string `form:"detail" binding:"-"`
}
