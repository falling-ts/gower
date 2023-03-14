package benchmarks

import "testing"

func BenchmarkConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cfg.App.Name
	}
}

func BenchmarkConfigGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cfg.Get("app.name").(string)
	}
}
