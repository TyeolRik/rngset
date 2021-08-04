package rngset

// Inversive Congruential Pseudorandom Numbers: A Tutorial
// Jiirgen Eichenauer-Herrman
// International Statistical Review (1992), 60, 2, pp. 167-176. Printed in Great Britain

func (r *icg) Generate() int64 {
	if r.seed == 0 {
		r.seed = r.c
	} else {
		x_inv := modularMultiplicativeInverse(r.seed, r.q)
		r.seed = (r.a*x_inv + r.c) % r.q
	}
	return r.seed
}

// Could be more speed with Euclidean algorithm
func modularMultiplicativeInverse(a, m int64) int64 {
	var i int64
	for i = 1; i < m; i++ {
		if (a*i)%m == 1 {
			return i
		}
	}
	return -1
}
