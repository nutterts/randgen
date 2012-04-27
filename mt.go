// Copyright 2012 Gert Nutterts. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package randgen

import "math/rand"

type MT struct {
	// Create two length 624 array to store the state of the generators
	arrayL, arrayU [624]uint32
	index, lp      uint32
}

var GlobalMT rand.Source

// New Mersenne Twister 
func NewMT() rand.Source {
	m := MT{index: 0}
	m.Seed(1)
	return &m
}

// Initialize the generator from a seed
func (m *MT) Seed(seed int64) {
	// Lower 32bit
	m.arrayL[0] = uint32(seed)
	for m.lp = 1; m.lp < 624; m.lp++ {
		m.arrayL[m.lp] = 0x6c078965*(m.arrayL[m.lp-1]^(m.arrayL[m.lp-1]>>30)) + m.lp
	}
	// Upper 32bit
	m.arrayU[0] = uint32(seed >> 32)
	for m.lp = 1; m.lp < 624; m.lp++ {
		m.arrayU[m.lp] = 0x6c078965*(m.arrayU[m.lp-1]^(m.arrayU[m.lp-1]>>30)) + m.lp
	}
}

// Extract a tempered pseudorandom number based on the index-th value,
// calling *MT.refill() every 624 numbers
func (m *MT) Int63() int64 {
	if m.index == 0 {
		m.refill()
	}
	// Lower bits
	yL := m.arrayL[m.index]
	yL ^= (yL >> 11)
	yL ^= ((yL << 7) & 0x9d2c5680)
	yL ^= ((yL << 15) & 0xefc60000)
	yL ^= (yL >> 18)
	// Upper bits
	yU := m.arrayU[m.index]
	yU ^= (yU >> 11)
	yU ^= ((yU << 7) & 0x9d2c5680)
	yU ^= ((yU << 15) & 0xefc60000)
	yU ^= (yU >> 18)

	m.index = (m.index + 1) % 624
	return (int64)((uint64)(yL) | ((uint64)(yU)<<32)&0x7FFFFFFFFFFFFFFF)
}

// Generate an array of 624 untempered numbers
func (m *MT) refill() {
	var y uint32
	for m.lp = 0; m.lp < 624; m.lp++ {
		// Lower bits
		y = (m.arrayL[m.lp] & 0x80000000) | (m.arrayL[(m.lp+1)%624] & 0xffffffff)
		m.arrayL[m.lp] = m.arrayL[(m.lp+397)%624] ^ (y >> 1)
		if (y % 2) != 0 { // y is odd			
			m.arrayL[m.lp] ^= 0x9908b0df
		}
		// Upper bits
		y = (m.arrayU[m.lp] & 0x80000000) | (m.arrayU[(m.lp+1)%624] & 0xffffffff)
		m.arrayU[m.lp] = m.arrayU[(m.lp+397)%624] ^ (y >> 1)
		if (y % 2) != 0 { // y is odd			
			m.arrayU[m.lp] ^= 0x9908b0df
		}
	}
}

func init() {
	GlobalMT = NewMT()
}
