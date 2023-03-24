package route

import (
	"gower/app"
	"gower/app/middlewares"
)

var route = app.Route()

func init() {
	route.Use(middlewares.Recovery()).
		Use(middlewares.Logger()).
		Use(middlewares.Cors()).
		Use(middlewares.CsrfToken())

	route.NoRoute([]any{
		"excp/404", app.Data{
			"msg":    "请求地址不存在",
			"detail": "非常抱歉，您所请求的页面或资源未找到。我们深表歉意，给您带来了不便。",
		},
	})
}
