// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import "math/rand"
import "testing"

func BenchmarkBuildin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int63()
	}
}
