// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import "math/rand"

type Kiss struct {
	x, c, y, z, t uint64
}

var GlobalKiss rand.Source

func NewKiss() rand.Source {
	return &Kiss{
		x: 1234567890987654321,
		c: 1,
		y: 362436362436362436,
		z: 1066149217761810,
	}
}

func (k *Kiss) Seed(seed int64) {
	k.c = (uint64)(seed) & 0x3FFFFFFFFFFFFFF
	k.x = 1234567890987654321
	k.y = 362436362436362436
	k.z = 1066149217761810
}

func (k *Kiss) Int63() int64 {
	// MWC
	k.t = (k.x << 58) + k.c
	k.c = (k.x >> 6)
	k.x += k.t
	if k.x < k.t {
		k.c++
	}
	// XSH
	k.y ^= k.y << 13
	k.y ^= k.y >> 17
	k.y ^= k.y << 43
	// CNG
	k.z = 6906969069*k.z + 1234567
	// Result
	return (int64)((k.x + k.y + k.z) & 0x7FFFFFFFFFFFFFFF)
}

func init() {
	GlobalKiss = NewKiss()
}
