package tests

import (
	"testing"
)

func TestConfig(t *testing.T) {
	assert := getAssert(t)
	assert.Equal(config.App.Name, config.Get("app.name").(string))
}
