package tests

import (
	"github.com/stretchr/testify/assert"
	"gower/app"
	"testing"
)

var (
	route = app.Route()
	cfg   = app.Configs()
	excp  = app.Exceptions()
	cache = app.Cache()
)

func getAssert(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
