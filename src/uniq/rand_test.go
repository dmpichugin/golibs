package uniq

import "testing"

func BenchmarkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int()
	}
}

func BenchmarkInt31(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int31()
	}
}

func BenchmarkInt31n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int31n(255)
	}
}

func BenchmarkInt63(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int63()
	}
}

func BenchmarkInt63n(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int63n(255)
	}
}
