// Improved Long-Period Generators Based on Linear Recurrences Modulo 2
// FRANCOIS PANNETON and PIERRE L’ECUYER
// Universite de Montreal

// http://www.iro.umontreal.ca/~panneton/well/WELL512a.c
// http://www.iro.umontreal.ca/~panneton/well/WELL1024a.c
// http://www.iro.umontreal.ca/~panneton/well/WELL19937a.c
// http://www.iro.umontreal.ca/~panneton/well/WELL44497a.c

package rngset

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
func (r *well512a) NextFloat64() float64 {
	r.z0 = r.state[(r.state_i+15)&0x0000000F]
	r.z1 = r.mat0neg(16, r.state[r.state_i]) ^ r.mat0neg(15, r.state[(r.state_i+r.m1)&0x0000000F])
	r.z2 = r.mat0pos(11, r.state[(r.state_i+r.m2)&0x0000000F])
	r.state[r.state_i] = r.z1 ^ r.z2
	r.state[(r.state_i+15)&0x0000000F] = r.mat0neg(2, r.z0) ^ r.mat0neg(18, r.z1) ^ r.mat3neg(28, r.z2) ^ r.mat4neg(5, 0xDA442D24, r.state[r.state_i])
	r.state_i = (r.state_i + 15) & 0x0000000F
	return float64(r.state[r.state_i]) * 2.32830643653869628906e-10
}

func (r *well1024a) mat0pos(t, v uint32) uint32 {
	return (v ^ (v >> t))
}
func (r *well1024a) mat0neg(t, v uint32) uint32 {
	return (v ^ (v << (t)))
}
func (r *well1024a) NextFloat64() float64 {
	r.z0 = r.state[(r.state_i+31)&0x0000001F]
	r.z1 = r.state[r.state_i] ^ r.mat0pos(8, r.state[(r.state_i+r.m1)&0x0000001F])
	r.z2 = r.mat0neg(19, r.state[(r.state_i+r.m2)&0x0000001F]) ^ r.mat0neg(14, r.state[(r.state_i+r.m3)&0x0000001F])
	r.state[r.state_i] = r.z1 ^ r.z2                                                                    // newV1
	r.state[(r.state_i+31)&0x0000001F] = r.mat0neg(11, r.z0) ^ r.mat0neg(7, r.z1) ^ r.mat0neg(13, r.z2) // newV0
	r.state_i = (r.state_i + 31) & 0x0000001F
	return float64(r.state[r.state_i]) * 2.32830643653869628906e-10
}

func (r *well19937a) mat0pos(t, v uint32) uint32 {
	return (v ^ (v >> t))
}
func (r *well19937a) mat0neg(t, v uint32) uint32 {
	return (v ^ (v << (t)))
}
func (r *well19937a) mat3pos(t, v uint32) uint32 {
	return (v >> t)
}
func (r *well19937a) case1() float64 {
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
		return float64(y) * 2.32830643653869628906e-10
	} else {
		return float64(r.state[r.state_i]) * 2.32830643653869628906e-10
	}
}
func NextFloat64() float64 {
	return 0
}
