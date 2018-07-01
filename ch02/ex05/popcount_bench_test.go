package ex05

import "testing"
import "../popcount"
import "../ex03"
import (
	"../ex04"
)

var output int

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output = popcount.PopCount(100)
	}
}

func BenchmarkPopCountEx3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output = ex03.PopCount(100)
	}
}

// full bit たってると64回
func BenchmarkPopCountEx4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output = ex04.PopCount(100)
	}
}

func BenchmarkPopCountEx5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		output = PopCount(100)
	}
}
