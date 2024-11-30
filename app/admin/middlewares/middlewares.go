package middlewares

import "gitee.com/falling-ts/gower/app"

var (
	db    = app.DB()
	trans = app.Translate()
	exc   = app.Exception()
)
