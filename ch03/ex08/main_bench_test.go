package main

import (
	"io/ioutil"
	"testing"
)

func BenchmarkFractalByComplex64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run("complex64", 100, 100, ioutil.Discard)
	}
}

func BenchmarkFractalByComplex128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run("complex128", 100, 100, ioutil.Discard)
	}
}

func BenchmarkFractalByFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run("float", 100, 100, ioutil.Discard)
	}
}

func BenchmarkFractalByRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run("rat", 100, 100, ioutil.Discard)
	}
}
