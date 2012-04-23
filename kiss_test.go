package randgen

import "testing"

func TestKiss(t *testing.T) {
	var r int64
	for i := 0; i < 100000000; i++ {
		r = GlobalKiss.Int63()
	}
	if r != 1666297717051644203 {
		t.Fail()
	}
}

func BenchmarkKiss(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GlobalKiss.Int63()
	}
}
