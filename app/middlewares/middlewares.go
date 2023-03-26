package middlewares

import (
	"github.com/falling-ts/gower/app"
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
