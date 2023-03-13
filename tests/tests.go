package tests

import (
	"github.com/stretchr/testify/assert"
	"gower/app"
	"testing"
)

var (
	App   = app.App
	cfg   = app.App.Configs()
	excp  = app.App.Exceptions()
	cache = App.Cache()
)

func getAssert(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
