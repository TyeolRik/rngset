package rngset

// ON THE LATTICE STRUCTURE OF THE
// ADD-WITH-CARRY AND SUBTRACT-WITH-BORROW
// RANDOM NUMBER GENERATORS
// SHU TEZUKA, IBM Research, Tokyo Research Laboratory
// PIERRE L’ECUYER, Universite de Montreal
// Page 73. 1. THE AWC AND SWB GENERATORS

func (r *awc) NextInt64() int64 {
	next := r.seed[0] + r.seed[r.r-r.s] + r.c
	if next >= r.b {
		r.c = 1
		next = next - r.b
	} else {
		r.c = 0
	}
	var i int64
	for i = 0; i < r.r-1; i++ {
		r.seed[i] = r.seed[i+1]
	}
	r.seed[r.r-1] = next
	return next
}

func (r *awc) NextFloat64() float64 {
	// According to equation (3)
	// If b is large enough (e.g., a large power of two), one can just take L=1
	return float64(r.NextInt64()) / float64(r.b)
}

func (r *awc_c) NextInt64() int64 {
	temp := r.seed[0] + r.seed[r.r-r.s] + r.c
	next := ((2*r.b - 1) - temp) % r.b
	if temp >= r.b {
		r.c = 1
	} else {
		r.c = 0
	}
	var i int64
	for i = 0; i < r.r-1; i++ {
		r.seed[i] = r.seed[i+1]
	}
	r.seed[r.r-1] = next
	return next
}

func (r *awc_c) NextFloat64() float64 {
	return float64(r.NextInt64()) / float64(r.b)
}

func (r *swb1) NextInt64() int64 {
	next := r.seed[r.r-r.s] - r.seed[0] - r.c
	if next < 0 {
		r.c = 1
		next = next + r.b
	} else {
		r.c = 0
	}
	var i int64
	for i = 0; i < r.r-1; i++ {
		r.seed[i] = r.seed[i+1]
	}
	r.seed[r.r-1] = next
	return next
}

func (r *swb1) NextFloat64() float64 {
	return float64(r.NextInt64()) / float64(r.b)
}

func (r *swb2) NextInt64() int64 {
	next := r.seed[0] - r.seed[r.r-r.s] - r.c
	if next < 0 {
		r.c = 1
		next = next + r.b
	} else {
		r.c = 0
	}
	var i int64
	for i = 0; i < r.r-1; i++ {
		r.seed[i] = r.seed[i+1]
	}
	r.seed[r.r-1] = next
	return next
}

func (r *swb2) NextFloat64() float64 {
	return float64(r.NextInt64()) / float64(r.b)
}
