package uniq

import "testing"

func BenchmarkUUIDv4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UUID()
	}
}
