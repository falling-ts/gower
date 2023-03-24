package middlewares

import (
	"gower/app"
)

var (
	config = app.Config()
	logger = app.Logger()
	excp   = app.Exception()
	auth   = app.Auth()
	db     = app.DB()
	trans  = app.Translate()
	cookie = app.Cookie()
)
