package middlewares

import "gower/app"

var (
	config = app.Config()
	logger = app.Logger()
	route  = app.Route()
	excp   = app.Exception()
	auth   = app.Auth()
	db     = app.DB()
	trans  = app.Translate()
)

func init() {
	route.Use(Recovery()).Use(Logger())
}
