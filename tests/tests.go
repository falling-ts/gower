package tests

import (
	"github.com/stretchr/testify/assert"
	"gower/app"
	"testing"
)

var (
	auth     = app.Auth()
	cache    = app.Cache()
	config   = app.Config()
	cookie   = app.Cookie()
	db       = app.DB()
	excp     = app.Exception()
	logger   = app.Logger()
	passwd   = app.Passwd()
	res      = app.Response()
	route    = app.Route()
	symCrypt = app.SymCrypt()
	trans    = app.Translate()
	util     = app.Util()
	valid    = app.Validator()
)

func getAssert(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
