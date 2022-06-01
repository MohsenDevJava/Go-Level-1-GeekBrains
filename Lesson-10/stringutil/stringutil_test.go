package stringutil

import "testing"

func TestStringUtil(t *testing.T) {
	SampleReverse()
}

func BenchmarkTestStringUtil(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SampleReverse()
	}
}
