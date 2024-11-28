package middlewares

import "gitee.com/falling-ts/gower/app"

var (
	db    = app.DB()
	trans = app.Translate()
	excp  = app.Exception()
)
