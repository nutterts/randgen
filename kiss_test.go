// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import "testing"

func TestKiss(t *testing.T) {
	var r int64
	for i := 0; i < 100000; i++ {
		r = GlobalKiss.Int63()
	}
	if r != 2239460713827829247 {
		t.Fail()
	} else {
		println("KISS [OK]")
	}
}

func BenchmarkKiss(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GlobalKiss.Int63()
	}
}
