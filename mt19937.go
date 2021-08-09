package rngset

import "log"

// https://en.wikipedia.org/wiki/Mersenne_Twister#Pseudocode

func (r *mt19937) NextUint32() uint32 {
	if r.index >= 624 {
		if r.index > 624 {
			log.Fatalln("Generator was never seeded")
		}
		// twist()
		var lowerMask uint32 = 0x7FFFFFFF // uint32((1 << 31) - 1)
		var upperMask uint32 = 0x80000000 // ^lowerMask
		for i := 0; i < 624; i++ {
			x := (r.MT[i] & upperMask) + (r.MT[(i+1)%624] & lowerMask)
			xA := x >> 1
			if x%2 != 0 {
				xA = xA ^ 0x9908B0DF
			}
			r.MT[i] = r.MT[(i+397)%624] ^ xA
		}
		r.index = 0
	}

	// (u, d) = (11, 0xFFFFFFFF)
	// (s, b) = ( 7, 0x9D2C5680)
	// (t, c) = (15, 0xEFC60000)
	// l = 18

	y := r.MT[r.index]
	y = y ^ ((y >> 11) & 0xFFFFFFFF)
	y = y ^ ((y << 7) & 0x9D2C5680)
	y = y ^ ((y << 15) & 0xEFC60000)
	y = y ^ (y >> 18)

	r.index += 1
	return y // lowest 32 bits of (y)
}

func (r *mt19937) NextFloat64() float64 {
	return float64(r.NextUint32()) / float64(0xFFFFFFFFFFFFFFFF)
}
