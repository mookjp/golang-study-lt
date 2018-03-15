package ex04

import "testing"
import "../popcount"
import "../ex03"

func BenchmarkPopCountOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(100)
	}
}

func BenchmarkPopCountEx3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ex03.PopCount(100)
	}
}

func BenchmarkPopCountEx4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(100)
	}
}
