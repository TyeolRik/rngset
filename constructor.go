package rngset

import (
	"fmt"
	"strconv"
)

type goMathRand struct {
	seed              int64
	isMathRandSeedSet bool
}

func NewGoMathRand(seed int64) goMathRand {
	return goMathRand{seed: seed, isMathRandSeedSet: false}
}

type goCryptoRand struct {
}

func NewGoCryptoRand() goCryptoRand {
	return goCryptoRand{}
}

type middleSquare struct {
	seed int64
}

func NewMiddleSquare(seed int64) middleSquare {
	return middleSquare{seed: seed}
}

type zx81 struct {
	seed int64
}

func NewZX81(seed int64) zx81 {
	return zx81{seed: seed}
}

type minstd_rand struct {
	seed int64
}

func NewMINSTD_RAND(seed int64) minstd_rand {
	return minstd_rand{seed: seed}
}

type wichmannHill struct {
	s1 int64
	s2 int64
	s3 int64
}

func NewWichmannHill(_s1, _s2, _s3 int64) wichmannHill {
	return wichmannHill{s1: _s1, s2: _s2, s3: _s3}
}

type rule30 struct {
	board  [][]uint8
	width  int64
	height int64
	seed   []uint8
}

func NewRule30(seed int64, height int64) rule30 {
	seed2bitstring := fmt.Sprintf("%b", seed)
	newRule := new(rule30)
	newRule.height = height
	var seedLength int64 = int64(len(seed2bitstring))
	if seedLength%2 == 0 {
		newRule.width = (newRule.height + (seedLength / 2)) * 2
		newRule.board = make([][]uint8, newRule.height)
		for i := range newRule.board {
			newRule.board[i] = make([]uint8, newRule.width)
		}
	} else {
		newRule.width = (newRule.height + (seedLength / 2)) * 2
		newRule.board = make([][]uint8, newRule.height)
		for i := range newRule.board {
			newRule.board[i] = make([]uint8, newRule.width)
		}
	}
	// Set Seed
	newRule.seed = make([]uint8, len(seed2bitstring))
	for i := range newRule.seed {
		bit, err := strconv.ParseUint(seed2bitstring[i:i+1], 10, 8)
		if err != nil {
			panic(err)
		}
		newRule.seed[i] = uint8(bit)
	}
	newRule.setSeed(newRule.seed)
	newRule.refresh()
	return *newRule
}

// nextSeed = (a * seed^(-1) + c) mod q
type icg struct {
	q    int64
	a    int64
	c    int64
	seed int64
}

// Inversive Congruential Generator
// https://en.wikipedia.org/wiki/Inversive_congruential_generator
func NewInversiveCongruentialGenerator(q, a, c, seed int64) icg {
	return icg{q: q, a: a, c: c, seed: seed}
}

// Additive COngruential Random Number
// Couldn't understand ACORN FORTRAN 77, http://acorn.wikramaratna.org/download.html
// So, I built on my own by reading paper
// The additive congruential random number generator A special case of a multiple recursive generator
// Roy S. Wikramaratna, Volume 216, Issue 2, 1 July 2008, Pages 371-387
type acorn struct {
	modulus int64
	order   int64 // k
	seed    []int64
}

func NewACORN(k int64) acorn {
	seeds := make([]int64, k+1)
	var mod int64 = 1 << 61
	r := NewGoCryptoRand()
	for i := range seeds {
		s := int64(float64(mod) * r.Float64(0.0, 1.0))
		if s%2 == 0 {
			s = s + 1
		}
		seeds[i] = s
	}
	return acorn{
		modulus: mod,
		order:   k,
		seed:    seeds,
	}
}
