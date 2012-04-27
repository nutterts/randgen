// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import "testing"

func TestMT(t *testing.T) {
	var r int64
	for i := 0; i < 100000; i++ {
		r = GlobalMT.Int63()
	}
	if r != 8026505133492131828 {
		t.Fail()
	} else {
		println("MT [OK]")
	}
}

func BenchmarkMT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GlobalMT.Int63()
	}
}
