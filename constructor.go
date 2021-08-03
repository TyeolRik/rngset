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

type icg struct {
	q    int64
	a    int64
	c    int64
	seed int64
}

// Inversive Congruential Generator
// https://en.wikipedia.org/wiki/Inversive_congruential_generator
func NewInversiveCongruentialGenerator(q, a, c, seed int64) icg {

}
