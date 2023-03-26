package main

import (
	"gower/tests"
	"gower/tests/benchmarks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("TestAuth", tests.TestAuth)
	t.Run("TestCache", tests.TestCache)
	t.Run("TestConfig", tests.TestConfig)
	t.Run("TestDB", tests.TestDB)
	t.Run("TestException", tests.TestException)

	t.Run("TextPasswd", tests.TextPasswd)
	t.Run("TestTrans", tests.TestTrans)

	t.Run("TestControllers", tests.TestControllers)
}

func Benchmark(b *testing.B) {
	b.Run("BenchmarkCache", benchmarks.BenchmarkCache)
	b.Run("BenchmarkConfig", benchmarks.BenchmarkConfig)
	b.Run("BenchmarkConfigGet", benchmarks.BenchmarkConfigGet)
}

func Example() {

}

func Fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		a := assert.New(t)
		a.IsType("string", s)
	})
}
