package randgen

import "testing"

func TestMT(t *testing.T) {
	var r int64
	for i := 0; i < 100000000; i++ {
		r = GlobalMT.Int63()
	}
	if r != 6696354919239096934 {
		t.Fail()
	}
}

func BenchmarkMT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GlobalMT.Int63()
	}
}