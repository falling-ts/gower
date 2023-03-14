package benchmarks

import "testing"

func BenchmarkCache(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache.SetDefault("key", "test")
		_, _ = cache.Get("key")
		cache.Delete("key")
	}
}
