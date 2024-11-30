package tests

import (
	"gitee.com/falling-ts/gower/app"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	auth     = app.Auth()
	cache    = app.Cache()
	config   = app.Config()
	db       = app.DB()
	exc      = app.Exception()
	passwd   = app.Passwd()
	res      = app.Response()
	route    = app.Route()
	symCrypt = app.SymCrypt()
	_        = app.Translate()
	util     = app.Util()
)

func getAssert(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
