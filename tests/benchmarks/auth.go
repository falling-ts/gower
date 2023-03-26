package benchmarks

import "testing"

func BenchmarkAuthSign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = auth.Sign("test_model", "5555", []string{"192.168.10.11:8080"})
	}
}

func BenchmarkAuthCheck(b *testing.B) {
	token, _ := auth.Sign("test_model", "5555", []string{"192.168.10.11:8080"})
	for i := 0; i < b.N; i++ {
		_, _, _ = auth.Check(token, "192.168.10.11:8080")
	}
}
