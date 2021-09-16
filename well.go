// Improved Long-Period Generators Based on Linear Recurrences Modulo 2
// FRANCOIS PANNETON and PIERRE L’ECUYER
// Universite de Montreal

// http://www.iro.umontreal.ca/~panneton/well/WELL512a.c
// http://www.iro.umontreal.ca/~panneton/well/WELL1024a.c
// http://www.iro.umontreal.ca/~panneton/well/WELL19937a.c
// http://www.iro.umontreal.ca/~panneton/well/WELL44497a.c

package rngset

import "log"

func (r *well512a) mat0pos(t, v uint32) uint32 {
	return (v ^ (v >> t))
}
func (r *well512a) mat0neg(t, v uint32) uint32 {
	return (v ^ (v << t))
}
func (r *well512a) mat3neg(t, v uint32) uint32 {
	return (v << t)
}
func (r *well512a) mat4neg(t, b, v uint32) uint32 {
	return (v ^ ((v << t) & b))
}
func (r *well512a) NewUint32() uint32 {
	r.z0 = r.state[(r.state_i+15)&0x0000000F]
	r.z1 = r.mat0neg(16, r.state[r.state_i]) ^ r.mat0neg(15, r.state[(r.state_i+r.m1)&0x0000000F])
	r.z2 = r.mat0pos(11, r.state[(r.state_i+r.m2)&0x0000000F])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[(r.state_i+15)&0x0000000F] = r.mat0neg(2, r.z0) ^ r.mat0neg(18, r.z1) ^ r.mat3neg(28, r.z2) ^ r.mat4neg(5, 0xDA442D24, r.state[r.state_i])
	r.state_i = (r.state_i + 15) & 0x0000000F
	return r.state[r.state_i]
}
func (r *well512a) NextFloat64() float64 {
	return float64(r.NewUint32()) * 2.32830643653869628906e-10
}

func (r *well1024a) mat0pos(t, v uint32) uint32 {
	return (v ^ (v >> t))
}
func (r *well1024a) mat0neg(t, v uint32) uint32 {
	return (v ^ (v << (t)))
}
func (r *well1024a) NextUint32() uint32 {
	r.z0 = r.state[(r.state_i+31)&0x0000001F]
	r.z1 = r.state[r.state_i] ^ r.mat0pos(8, r.state[(r.state_i+r.m1)&0x0000001F])
	r.z2 = r.mat0neg(19, r.state[(r.state_i+r.m2)&0x0000001F]) ^ r.mat0neg(14, r.state[(r.state_i+r.m3)&0x0000001F])
	r.state[r.state_i] = r.z1 ^ r.z2                                                                    // r.state[r.state_i]
	r.state[(r.state_i+31)&0x0000001F] = r.mat0neg(11, r.z0) ^ r.mat0neg(7, r.z1) ^ r.mat0neg(13, r.z2) // r.state[r.state_i-1]
	r.state_i = (r.state_i + 31) & 0x0000001F
	return r.state[r.state_i]
}
func (r *well1024a) NextFloat64() float64 {
	return float64(r.NextUint32()) * 2.32830643653869628906e-10
}

func (r *well19937) mat0pos(t, v uint32) uint32 {
	return (v ^ (v >> t))
}
func (r *well19937) mat0neg(t, v uint32) uint32 {
	return (v ^ (v << (t)))
}
func (r *well19937) mat3pos(t, v uint32) uint32 {
	return (v >> t)
}
func (r *well19937) case1() uint32 {
	// r.state_i == 0
	r.z0 = (r.state[r.state_i+r.r-1] & r.maskL) | (r.state[r.state_i+r.r-2] & r.maskU)
	r.z1 = r.mat0neg(25, r.state[r.state_i]) ^ r.mat0pos(27, r.state[r.state_i+r.m1])
	r.z2 = r.mat3pos(9, r.state[r.state_i+r.m2]) ^ r.mat0pos(1, r.state[r.state_i+r.m3])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[r.state_i-1+r.r] = r.z0 ^ r.mat0neg(9, r.z1) ^ r.mat0neg(21, r.z2) ^ r.mat0pos(21, r.state[r.state_i])
	r.state_i = r.r - 1
	r._case = 3

	if r.tempering {
		y := r.state[r.state_i] ^ ((r.state[r.state_i] << 7) & r.temperB)
		y = y ^ ((y << 15) & r.temperC)
		return y
	} else {
		return r.state[r.state_i]
	}
}
func (r *well19937) case2() uint32 {
	// r.state_i == 1
	r.z0 = (r.state[r.state_i-1] & r.maskL) | (r.state[r.state_i+r.r-2] & r.maskU)
	r.z1 = r.mat0neg(25, r.state[r.state_i]) ^ r.mat0pos(27, r.state[r.state_i+r.m1])
	r.z2 = r.mat3pos(9, r.state[r.state_i+r.m2]) ^ r.mat0pos(1, r.state[r.state_i]+r.m3)
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[r.state_i-1] = r.z0 ^ r.mat0neg(9, r.z1) ^ r.mat0neg(21, r.z2) ^ r.mat0pos(21, r.state[r.state_i])
	r.state_i = 0
	r._case = 1
	if r.tempering {
		y := r.state[r.state_i] ^ ((r.state[r.state_i] << 7) & r.temperB)
		y = y ^ ((y << 15) & r.temperC)
		return y
	} else {
		return r.state[r.state_i]
	}
}
func (r *well19937) case3() uint32 {
	// r.state_i+r.m1 >= r.r
	r.z0 = (r.state[r.state_i-1] & r.maskL) | (r.state[r.state_i-2] & r.maskU)
	r.z1 = r.mat0neg(25, r.state[r.state_i]) ^ r.mat0pos(27, r.state[r.state_i+r.m1-r.r])
	r.z2 = r.mat3pos(9, r.state[r.state_i+r.m2-r.r]) ^ r.mat0pos(1, r.state[r.state_i+r.m3-r.r])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[r.state_i-1] = r.z0 ^ r.mat0neg(9, r.z1) ^ r.mat0neg(21, r.z2) ^ r.mat0pos(21, r.state[r.state_i])
	r.state_i--
	if r.state_i+r.m1 < r.r {
		r._case = 5
	}
	if r.tempering {
		y := r.state[r.state_i] ^ ((r.state[r.state_i] << 7) & r.temperB)
		y = y ^ ((y << 15) & r.temperC)
		return y
	} else {
		return r.state[r.state_i]
	}
}
func (r *well19937) case4() uint32 {
	// r.state_i+r.m3 >= r.r
	r.z0 = (r.state[r.state_i-1] & r.maskL) | (r.state[r.state_i-2] & r.maskU)
	r.z1 = r.mat0neg(25, r.state[r.state_i]) ^ r.mat0pos(27, r.state[r.state_i+r.m1])
	r.z2 = r.mat3pos(9, r.state[r.state_i+r.m2]) ^ r.mat0pos(1, r.state[r.state_i+r.m3-r.r])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[r.state_i-1] = r.z0 ^ r.mat0neg(9, r.z1) ^ r.mat0neg(21, r.z2) ^ r.mat0pos(21, r.state[r.state_i])
	r.state_i--
	if r.state_i+r.m3 < r.r {
		r._case = 6
	}
	if r.tempering {
		y := r.state[r.state_i] ^ ((r.state[r.state_i] << 7) & r.temperB)
		y = y ^ ((y << 15) & r.temperC)
		return y
	} else {
		return r.state[r.state_i]
	}
}
func (r *well19937) case5() uint32 {
	// r.state_i+r.m2 >= r.r
	r.z0 = (r.state[r.state_i-1] & r.maskL) | (r.state[r.state_i-2] & r.maskU)
	r.z1 = r.mat0neg(25, r.state[r.state_i]) ^ r.mat0pos(27, r.state[r.state_i+r.m1])
	r.z2 = r.mat3pos(9, r.state[r.state_i+r.m2-r.r]) ^ r.mat0pos(1, r.state[r.state_i+r.m3-r.r])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[r.state_i-1] = r.z0 ^ r.mat0neg(9, r.z1) ^ r.mat0neg(21, r.z2) ^ r.mat0pos(21, r.state[r.state_i])
	r.state_i--
	if r.state_i+r.m2 < r.r {
		r._case = 4
	}
	if r.tempering {
		y := r.state[r.state_i] ^ ((r.state[r.state_i] << 7) & r.temperB)
		y = y ^ ((y << 15) & r.temperC)
		return y
	} else {
		return r.state[r.state_i]
	}
}
func (r *well19937) case6() uint32 {
	// 2 <= r.state_i <= (r.r - r.m3 - 1)
	r.z0 = (r.state[r.state_i-1] & r.maskL) | (r.state[r.state_i-2] & r.maskU)
	r.z1 = r.mat0neg(25, r.state[r.state_i]) ^ r.mat0pos(27, r.state[r.state_i+r.m1])
	r.z2 = r.mat3pos(9, r.state[r.state_i+r.m2]) ^ r.mat0pos(1, r.state[r.state_i+r.m3])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[r.state_i-1] = r.z0 ^ r.mat0neg(9, r.z1) ^ r.mat0neg(21, r.z2) ^ r.mat0pos(21, r.state[r.state_i])
	r.state_i--
	if r.state_i == 1 {
		r._case = 2
	}
	if r.tempering {
		y := r.state[r.state_i] ^ ((r.state[r.state_i] << 7) & r.temperB)
		y = y ^ ((y << 15) & r.temperC)
		return y
	} else {
		return r.state[r.state_i]
	}
}
func (r *well19937) NewUint32() uint32 {
	switch r._case {
	case 1:
		return r.case1()
	case 2:
		return r.case2()
	case 3:
		return r.case3()
	case 4:
		return r.case4()
	case 5:
		return r.case5()
	case 6:
		return r.case6()
	default:
		log.Fatalln("WELL19937a :: Something is wrong!!")
		return 0
	}
}
func (r *well19937) NextFloat64() float64 {
	return float64(r.NewUint32()) * 2.32830643653869628906e-10
}
