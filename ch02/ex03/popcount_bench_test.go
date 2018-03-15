package ex03

import "testing"
import "../popcount"

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100)
	}
}

func BenchmarkPopCountEx3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(100)
	}
}
