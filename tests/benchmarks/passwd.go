package benchmarks

import "testing"

func BenchmarkPasswd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = passwd.Hash("125455885fds54fd5s")
	}
}
