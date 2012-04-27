package randgen

import "math/rand"
import "testing"

func BenchmarkBuildin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int63()
	}
}
