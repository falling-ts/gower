package tests

import (
	"github.com/stretchr/testify/assert"
	"gower/app"
	"testing"
)

var (
	App    = app.Entity
	cfg    = app.Cfg()
	config = App.Config()
	excp   = app.Excp()
	cache  = App.Cache()
)

func getAssert(t *testing.T) *assert.Assertions {
	return assert.New(t)
}
