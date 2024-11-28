package main

import (
	"gitee.com/falling-ts/gower/tests"
	"gitee.com/falling-ts/gower/tests/benchmarks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("TestAuth", tests.TestAuth)
	t.Run("TestCache", tests.TestCache)
	t.Run("TestConfig", tests.TestConfig)
	t.Run("TestDB", tests.TestDB)
	t.Run("TestException", tests.TestException)
	t.Run("TestPasswd", tests.TextPasswd)
	t.Run("TestRoute", tests.TestRoute)
	t.Run("TestSymCrypt", tests.TestSymCrypt)
	t.Run("TestTrans", tests.TestTrans)
	t.Run("TestUtil", tests.TestUtil)
}

func Benchmark(b *testing.B) {
	b.Run("BenchmarkRoute01", benchmarks.BenchmarkRoute01)
	b.Run("BenchmarkRoute02", benchmarks.BenchmarkRoute02)
	b.Run("BenchmarkRoute03", benchmarks.BenchmarkRoute03)
	b.Run("BenchmarkRoute04", benchmarks.BenchmarkRoute04)
	b.Run("BenchmarkRoute05", benchmarks.BenchmarkRoute05)
	b.Run("BenchmarkRoute06", benchmarks.BenchmarkRoute06)
	b.Run("BenchmarkRoute07", benchmarks.BenchmarkRoute07)
	b.Run("BenchmarkRoute08", benchmarks.BenchmarkRoute08)
	b.Run("BenchmarkRoute09", benchmarks.BenchmarkRoute09)
	b.Run("BenchmarkRoute10", benchmarks.BenchmarkRoute10)
	b.Run("BenchmarkAuthSign", benchmarks.BenchmarkAuthSign)
	b.Run("BenchmarkAuthCheck", benchmarks.BenchmarkAuthCheck)
	b.Run("BenchmarkCache", benchmarks.BenchmarkCache)
	b.Run("BenchmarkConfig", benchmarks.BenchmarkConfig)
	b.Run("BenchmarkConfigGet", benchmarks.BenchmarkConfigGet)
	b.Run("BenchmarkPasswd", benchmarks.BenchmarkPasswd)
	b.Run("BenchmarkEncode", benchmarks.BenchmarkEncode)
	b.Run("BenchmarkDecode", benchmarks.BenchmarkDecode)
}

func Example() {}

func Fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		a := assert.New(t)
		a.IsType("string", s)
	})
}
