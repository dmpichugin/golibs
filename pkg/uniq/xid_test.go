package uniq

import "testing"

func BenchmarkXid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Xid()
	}
}
