package tests

import "testing"

func TestCache(t *testing.T) {
	assert := getAssert(t)
	cache.SetDefault("test-key", "123")
	value, ok := cache.Get("test-key")
	if !ok {
		t.Error("缓存错误")
	}
	assert.Equal(value, "123")
}
