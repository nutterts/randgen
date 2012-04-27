// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import "testing"

func TestIOReader(t *testing.T) {
	rng := NewKiss()
	t1 := (byte)(rng.Int63())

	io := Reader(NewKiss())
	var b [8]byte
	io.Read(b[:])
	t2 := b[7]

	if t1 != t2 {
		t.Fail()
	} else {
		println("IOReader [OK]")
	}
}
