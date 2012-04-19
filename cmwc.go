/*
	Copyleft 2012 Gert Nutterts, All rights reserved


*/
package randgen

const phi uint32 = 0x9e3779b9
const cs uint64 = 0x587C4

type CMWC struct {
	q [4096]uint32
	c uint64
}

func NewCMWC() *CMWC {
	c := CMWC{c: cs}
	c.Seed(1)
	return &c
}

func (c *CMWC) Seed32(seed uint32) {
	c.q[0] = seed
	c.q[1] = seed + phi
	c.q[2] = seed + phi + phi

	var i uint32
	for i = 3; i < 4096; i++ {
		c.q[i] = c.q[i-3] ^ c.q[i-2] ^ phi ^ i
	}
}

// math/rand.Source interface - Initialize the generator with a 32bit seed as an int64.
func (m *CMWC) Seed(seed int64) { m.Seed32((uint32)(seed)) }

func (c *CMWC) Int32() uint32 {
	var (
		i, x, r uint64 = 4095, 0xfffffffe, 0xfffffffe
		t, a    uint64 = 0, 18782
	)
	i = (i + 1) & 4095
	t = a*(uint64)(c.q[i]) + c.c
	c.c = t >> 32
	x = t + c.c
	if x < c.c {
		x++
		c.c++
	}
	c.q[i] = (uint32)(r - x)
	return c.q[i]
}

// math/rand.Source interface - return a pseudo-random 32-bit integer as an int64
func (m *CMWC) Int63() int64 { return (int64)(m.Int32()) }
