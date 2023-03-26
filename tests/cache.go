package tests

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	fmt.Println("----------------TestCache 开始----------------")
	
	assert := getAssert(t)

	cache.SetDefault("test-key", "123")
	value, ok := cache.Get("test-key")

	assert.True(ok)
	assert.Equal(value, "123")

	cache.Delete("test-key")

	_, ok = cache.Get("test-key")
	assert.False(ok)

	cache.Set("test-key", "123", time.Millisecond)
	time.Sleep(time.Millisecond * 10)

	_, ok = cache.Get("test-key")
	assert.False(ok)

	fmt.Println("----------------TestCache 结束----------------")
}
