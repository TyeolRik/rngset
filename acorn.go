package rngset

func (r *acorn) NextInt64() int64 {
	var i int64
	for i = 1; i <= r.order; i++ {
		r.seed[i] = (r.seed[i-1] + r.seed[i]) % r.modulus
	}
	return r.seed[r.order]
}

func (r *acorn) NextFloat64() float64 {
	return float64(r.NextInt64()) / float64(r.modulus)
}
