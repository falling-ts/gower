package benchmarks

import "testing"

func BenchmarkConfig(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = config.App.Name
	}
}

func BenchmarkConfigGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = config.Get("app.name").(string)
	}
}
