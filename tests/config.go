package tests

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	fmt.Println("----------------TestConfig 开始----------------")

	assert := getAssert(t)
	assert.Equal(config.App.Name, config.Get("app.name").(string))

	fmt.Println("----------------TestConfig 结束----------------")
}
