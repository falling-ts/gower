package routes

import (
	"github.com/falling-ts/gower/app"
	mws "github.com/falling-ts/gower/app/middlewares"
)

var route = app.Route()

func init() {
	route.Use(mws.Recovery()).
		Use(mws.Logger()).
		Use(mws.Cors()).
		Use(mws.CsrfToken())

	route.NoRoute([]any{
		"excp/404", app.Data{
			"msg":    "请求地址不存在",
			"detail": "非常抱歉，您所请求的页面或资源未找到。我们深表歉意，给您带来了不便。",
		},
	})
}
