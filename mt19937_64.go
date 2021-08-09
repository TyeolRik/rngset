package rngset

import "log"

// https://en.wikipedia.org/wiki/Mersenne_Twister#Pseudocode

func (r *mt19937_64) NextUint64() uint64 {
	if r.index >= 312 {
		if r.index > 312 {
			log.Fatalln("Generator was never seeded")
		}
		// twist()
		var lowerMask uint64 = 0x000000007FFFFFFF // uint64((1 << 31) - 1)
		var upperMask uint64 = 0xFFFFFFFF80000000 // ^lowerMask
		for i := 0; i < 312; i++ {
			x := (r.MT[i] & upperMask) + (r.MT[(i+1)%312] & lowerMask)
			xA := x >> 1
			if x%2 != 0 {
				xA = xA ^ 0xB5026F5AA96619E9
			}
			r.MT[i] = r.MT[(i+156)%312] ^ xA
		}
		r.index = 0
	}

	// (u, d) = (29, 0x5555555555555555)
	// (s, b) = (17, 0x71D67FFFEDA60000)
	// (t, c) = (37, 0xFFF7EEE000000000)
	// l = 43

	y := r.MT[r.index]
	y = y ^ ((y >> 29) & 0x5555555555555555)
	y = y ^ ((y << 17) & 0x71D67FFFEDA60000)
	y = y ^ ((y << 37) & 0xFFF7EEE000000000)
	y = y ^ (y >> 43)

	r.index += 1
	return y // lowest 64 bits of (y)
}

func (r *mt19937_64) NextFloat64() float64 {
	return float64(r.NextUint64()) / float64(0xFFFFFFFFFFFFFFFF)
}
