// Porting C to Golang.
// TyeolRik

// Keep it Simple Stupid, 64-bit MWC version (2011 version)
// Postname : RNGs with periods exceeding 10^(40million).
// https://www.thecodingforums.com/threads/rngs-with-periods-exceeding-10-40million.742134/

package rngset

func (r *kiss) b64mwc() uint64 {
	var t, x uint64
	r.j = (r.j + 1) & 2097151
	x = r.Q[r.j]
	t = (x << 28) + r.carry
	if t < x {
		r.carry = (x >> 36) - 1
	} else {
		r.carry = (x >> 36)
	}
	r.Q[r.j] = t - x
	return r.Q[r.j]
}

func (r *kiss) cng() uint64 {
	r._cng = uint64(6906969069)*(r._cng) + 13579
	return r._cng
}

func (r *kiss) xs() uint64 {
	r._xs ^= (r._xs << 13)
	r._xs ^= (r._xs >> 17)
	r._xs ^= (r._xs << 43)
	return r._xs
}

func (r *kiss) kiss() uint64 {
	return r.b64mwc() + r.cng() + r.xs()
}

func (r *kiss) NextUInt64() uint64 {
	return r.b64mwc() + r.cng() + r.xs()
}

func (r *kiss) NewFloat64() float64 {
	return float64(r.NextUInt64()) / float64(^uint64(0))
}
