package main

import (
	"gower/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	t.Run("TestConfig", tests.TestConfig)
	t.Run("TestException", tests.TestException)
	t.Run("TestCache", tests.TestCache)
}

func Benchmark(b *testing.B) {

}

func Example() {

}

func Fuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		a := assert.New(t)
		a.IsType("string", s)
	})
}
