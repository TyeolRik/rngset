package rngset

// Linear congruential generator
func LCG(a, xn, c, m int64) (x_next int64) {
	x_next = (a*xn + c) % m
	return
}

// https://en.wikipedia.org/wiki/Linear_congruential_generator#Parameters_in_common_use
func (r *zx81) Generate() int64 {
	r.seed = LCG(75, r.seed, 74, (1<<16)+1)
	return r.seed
}

func (r *minstd_rand) Generate() int64 {
	r.seed = LCG(48271, r.seed, 0, (1<<31)-1)
	if r.seed == 0 {
		r.seed = LCG(48271, 1, 0, (1<<31)-1)
	}
	return r.seed
}

// Wichmann–Hill generator
// Return [0, 1)
// B. A. WICHMAN, I. D. HILL, An Efficient and Portable Pseudo-random Number Generator, 1982, p.190
// doi:10.2307/2347988
func (r *wichmannHill) Generate() float64 {
	r.s1 = (171 * r.s1) % 30269
	r.s2 = (172 * r.s2) % 30307
	r.s3 = (170 * r.s3) % 30323
	ret := float64(r.s1)/30269.0 + float64(r.s2)/30307.0 + float64(r.s3)/30323.0
	return ret - float64(int64(ret))
}
