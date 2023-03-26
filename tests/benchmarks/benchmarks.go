package benchmarks

import (
	"gower/app"
)

var (
	auth     = app.Auth()
	cache    = app.Cache()
	config   = app.Config()
	passwd   = app.Passwd()
	route    = app.Route()
	symCrypt = app.SymCrypt()
	res      = app.Response()
)
