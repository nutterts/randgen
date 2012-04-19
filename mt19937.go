/*
	Copyleft 2012 Gert Nutterts, All rights reserved

	Based on the psuedocode version on the wikipedia page at:
	http://en.wikipedia.org/wiki/Mersenne_twister#Pseudocode
*/
package randgen

type MT19937 struct {
	// Create a length 624 array to store the state of the generator
	array [624]uint32
	index uint32 // index of array
}

// New Mersenne Twister 
func NewMT() *MT19937 {
	mt := MT19937{index: 0}
	mt.Seed(1)
	return &mt
}

// Initialize the generator from a 32bit seed
func (m *MT19937) Seed32(seed uint32) {
	m.array[0] = seed

	var i uint32
	for i = 1; i < 624; i++ {
		m.array[i] = 0x6c078965*(m.array[i-1]^(m.array[i-1]>>30)) + i
	}
}

// math/rand.Source interface - Initialize the generator with a 32bit seed as an int64.
func (m *MT19937) Seed(seed int64) { m.Seed32((uint32)(seed)) }

// Extract a tempered pseudorandom number based on the index-th value,
// calling *MT19937.generate() every 624 numbers
func (m *MT19937) Int32() uint32 {
	if m.index == 0 {
		m.generate()
	}

	y := m.array[m.index]
	y ^= (y >> 11)
	y ^= ((y << 7) & 0x9d2c5680)
	y ^= ((y << 15) & 0xefc60000)
	y ^= (y >> 18)

	m.index = (m.index + 1) % 624
	return y
}

// math/rand.Source interface - return a pseudo-random 32-bit integer as an int64
func (m *MT19937) Int63() int64 { return (int64)(m.Int32()) }

// Generate an array of 624 untempered numbers
func (m *MT19937) generate() {
	var i uint32
	for i = 0; i < 624; i++ {
		y := (m.array[i] & 0x80000000) | (m.array[(i+1)%624] & 0xffffffff)
		m.array[i] = m.array[(i+397)%624] ^ (y >> 1)
		if (y % 2) != 0 { // y is odd			
			m.array[i] ^= 0x9908b0df
		}
	}
}
