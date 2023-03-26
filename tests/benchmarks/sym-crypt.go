package benchmarks

import "testing"

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = symCrypt.Encrypt("14s5f1e56a1f5e6a15f6ae")
	}
}

func BenchmarkDecode(b *testing.B) {
	cipher, _ := symCrypt.Encrypt("14s5f1e56a1f5e6a15f6ae")
	for i := 0; i < b.N; i++ {
		_, _ = symCrypt.Decrypt(cipher)
	}
}
