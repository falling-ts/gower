package tests

import (
	"github.com/stretchr/testify/assert"
	"gower/app"
	"testing"
)

var (
	route = app.Route()
	cfg   = app.Config()
	excp  = app.Exception()
	cache = app.Cache()
	res   = app.Response()
)

func getAssert(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
