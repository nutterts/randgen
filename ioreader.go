// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import (
	"io"
	"math/rand"
)

type ioReader struct {
	s rand.Source
}

func Reader(r rand.Source) io.Reader {
	return &ioReader{s: r}
}

func (io *ioReader) Read(p []byte) (n int, err error) {
	c := len(p) / 8
	if c == 0 {
		return 0, nil
	}
	for lp := 0; lp < c; lp++ {
		i := lp * 8
		r := (uint64)(io.s.Int63())
		p[i] = (byte)(r)
		p[i+1] = (byte)(r >> 8)
		p[i+2] = (byte)(r >> 16)
		p[i+3] = (byte)(r >> 24)
		p[i+4] = (byte)(r >> 32)
		p[i+5] = (byte)(r >> 40)
		p[i+6] = (byte)(r >> 48)
		p[i+7] = (byte)(r >> 56)
	}
	return c * 8, nil
}
