package ex05

import "testing"
import "../popcount"
import "../ex03"
import "../ex04"

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
		ex04.PopCount(100)
	}
}

func BenchmarkPopCountEx5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(100)
	}
}
