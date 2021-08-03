package rngset

import (
	mathRand "math/rand"
	"strconv"
)

type randomGenerator struct {
	seed int64
}

func NewRNG(seed int64) (r randomGenerator) {
	r.seed = seed
	return
}

// Default math/rand in Go-language
// return [min, max)
func (r *goMathRand) Int64(min, max int64) int64 {
	if !r.isMathRandSeedSet {
		mathRand.Seed(r.seed)
		r.isMathRandSeedSet = true
	}
	return mathRand.Int63n(max-min) + min
}

// Default CSPRNG in Go-language
// return [min, max)
func (r *goCryptoRand) Int64(min, max int64) int64 {
	return GetCryptoRandInt(min, max)
}

// Middle-square method, John von Neumann (1949)
// https://en.wikipedia.org/wiki/Middle-square_method#Example_implementation
// Return random number with seed's digit.
// If Seed is 4 digit, returns 4 digit number.
func (r *middleSquare) Generate() int64 {
	// The value of n must be even in order for the method to work
	// if the value of n is odd then there will not necessarily be a uniquely defined 'middle n-digits' to select from.
	n := len(strconv.FormatInt(r.seed, 10))
	if n%2 == 1 {
		n = n + 1
	}
	squared := strconv.FormatInt(r.seed*r.seed, 10)
	for len(squared) < 2*n {
		squared = "0" + squared
	}
	newSeed, err := strconv.ParseInt(squared[len(squared)/2-n/2:len(squared)/2+n/2], 10, 64)
	if err != nil {
		panic(err)
	}
	r.seed = newSeed
	return newSeed
}
