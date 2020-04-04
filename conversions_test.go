package pipescript

import "testing"

func BenchmarkFloat(b *testing.B) {
	v := float64(3454.4)
	for n := 0; n < b.N; n++ {
		Float(v)
	}
}
