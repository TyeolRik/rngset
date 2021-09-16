package rngset

import (
	"fmt"
	"log"
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

// Add With Carry
type awc struct {
	seed []int64
	c    int64 // Carry

	b int64 // base, better to pick Prime Number
	r int64 // r > s
	s int64 // r > s
}

// Recommended suitable parameters
// According to Journal,
// George Marsaglia, Arif Zaman, A New Class of Random Number Generators, Page 476
// https://projecteuclid.org/journals/annals-of-applied-probability/volume-1/issue-3/A-New-Class-of-Random-Number-Generators/10.1214/aoap/1177005878.full
func NewAWC_Recommend() awc {
	// Couldn't understand Section 5.
	// Which is about seed.
	// So I used just random bits from Go default crypto/rand
	init := make([]int64, 39)
	r := NewGoCryptoRand()
	for i := range init {
		init[i] = r.Int64(0, 16777215) // On the p.473, each seed values in range 0 to (b - 1) (=2^24-1)
	}
	return awc{
		seed: init,
		c:    0,
		r:    39,
		s:    25,
		b:    (1 << 25),
	}
}

func NewAWC(seed []int64, r int64, s int64, b int64) awc {
	if r <= s {
		log.Fatalln("AWC :: ERROR : r should be greater than s!")
	}
	return awc{
		seed: seed,
		c:    0,
		r:    r,
		s:    s,
		b:    b,
	}
}

// The Complementary AWC generator
type awc_c struct {
	seed []int64
	c    int64 // Carry

	b int64 // base, better to pick Prime Number
	r int64 // r > s
	s int64 // r > s
}

func NewAWC_C_Recommend() awc_c {
	init := make([]int64, 39)
	r := NewGoCryptoRand()
	for i := range init {
		init[i] = r.Int64(0, 16777215) // On the p.473, each seed values in range 0 to (b - 1) (=2^24-1)
	}
	return awc_c{
		seed: init,
		c:    0,
		r:    39,
		s:    25,
		b:    (1 << 25),
	}
}

func NewAWC_C(seed []int64, r int64, s int64, b int64) awc_c {
	if r <= s {
		log.Fatalln("AWC_C :: ERROR : r should be greater than s!")
	}
	return awc_c{
		seed: seed,
		c:    0,
		r:    r,
		s:    s,
		b:    b,
	}
}

type swb1 struct {
	seed []int64
	c    int64 // Carry

	b int64 // base, better to pick Prime Number
	r int64 // r > s
	s int64 // r > s
}

func NewSWB1_Recommend() swb1 {
	init := make([]int64, 39)
	r := NewGoCryptoRand()
	for i := range init {
		init[i] = r.Int64(0, 16777215) // On the p.473, each seed values in range 0 to (b - 1) (=2^24-1)
	}
	return swb1{
		seed: init,
		c:    0,
		r:    39,
		s:    25,
		b:    (1 << 25),
	}
}

func NewSWB1(seed []int64, r int64, s int64, b int64) swb1 {
	if r <= s {
		log.Fatalln("AWC_C :: ERROR : r should be greater than s!")
	}
	return swb1{
		seed: seed,
		c:    0,
		r:    r,
		s:    s,
		b:    b,
	}
}

type swb2 struct {
	seed []int64
	c    int64 // Carry

	b int64 // base, better to pick Prime Number
	r int64 // r > s
	s int64 // r > s
}

func NewSWB2_Recommend() swb2 {
	init := make([]int64, 39)
	r := NewGoCryptoRand()
	for i := range init {
		init[i] = r.Int64(0, 16777215) // On the p.473, each seed values in range 0 to (b - 1) (=2^24-1)
	}
	return swb2{
		seed: init,
		c:    0,
		r:    39,
		s:    25,
		b:    (1 << 25),
	}
}

func NewSWB2(seed []int64, r int64, s int64, b int64) swb2 {
	if r <= s {
		log.Fatalln("AWC_C :: ERROR : r should be greater than s!")
	}
	return swb2{
		seed: seed,
		c:    0,
		r:    r,
		s:    s,
		b:    b,
	}
}

// Keep it Simple Stupid, 64-bit MWC version (2011 version)
// https://www.thecodingforums.com/threads/rngs-with-periods-exceeding-10-40million.742134/
type kiss struct {
	Q     [2097152]uint64
	carry uint64

	j    int32
	_cng uint64
	_xs  uint64
}

func NewKISS(cng, xs uint64) kiss {
	// First seed Q[] with CNG+XS:
	r := kiss{
		carry: 0,
		j:     2097151,
		_cng:  cng,
		_xs:   xs,
	}
	/*
		for i := range r.Q {
			// r.Q[i] = r.cng() + r.xs()  // For reducing function overhead
			r._xs ^= (r._xs << 13)
			r._xs ^= (r._xs >> 17)
			r._xs ^= (r._xs << 43)
			r._cng = uint64(6906969069)*r._cng + 13579
			r.Q[i] = r._cng + r._xs
		}
		// generate B64MWC()
		for i := 0; i < 1000000000; i++ {
			r.j = (r.j + 1) & 2097151
			x := r.Q[r.j]
			t := (x << 28) + r.carry
			if t < x {
				r.carry = (x >> 36) - 1
			} else {
				r.carry = (x >> 36)
			}
			r.Q[r.j] = t - x
		}
	*/
	return r
}

type mt19937 struct {
	MT    [624]uint32
	index uint32
}

func NewMT19937(seed uint32) mt19937 {
	r := mt19937{}
	r.index = 624
	r.MT[0] = seed
	var i uint32
	for i = 1; i < 624; i++ {
		// The value for f for MT19937 is 1812433253
		r.MT[i] = 1812433253*(r.MT[i-1]^(r.MT[i-1]>>30)) + i // lowest 32 bits
	}
	return r
}

type mt19937_64 struct {
	MT    [312]uint64
	index uint64
}

// Initialize the generator from a seed
func NewMT19937_64(seed uint64) mt19937_64 {
	r := mt19937_64{}
	r.index = 312
	r.MT[0] = seed
	var i uint64
	for i = 1; i < 312; i++ {
		// MT[i] := lowest w bits of (f * (MT[i-1] xor (MT[i-1] >> (w-2))) + i)
		// The value for f for MT19937-64 is 6364136223846793005.
		r.MT[i] = 6364136223846793005*(r.MT[i-1]^(r.MT[i-1]>>62)) + i // lowest 64 bits
	}
	return r
}

// Well equidistributed long-period linear
type well512a struct {
	state   [16]uint32
	state_i uint32

	w, r, p    uint32
	m1, m2, m3 uint32
	z0, z1, z2 uint32
}

// the length of Seeds should be 16
func NewWELL512a(seeds [16]uint32) (r well512a) {
	r = well512a{
		state_i: 0,
		w:       32,
		r:       16,
		p:       0,
		m1:      13,
		m2:      9,
		m3:      5,
	}
	copy(r.state[:], seeds[:])
	return
}

type well1024a struct {
	state   [32]uint32
	state_i uint32

	w, r       uint32
	m1, m2, m3 uint32
	z0, z1, z2 uint32
}

func NewWELL1024a(seeds [32]uint32) (r well1024a) {
	r = well1024a{
		state_i: 0,
		w:       32,
		r:       32,
		m1:      3,
		m2:      24,
		m3:      10,
	}
	copy(r.state[:], seeds[:])
	return
}

type well19937 struct {
	state   [624]uint32
	state_i uint32

	_case            uint32 // case 1, 2, 3, 4, 5, 6
	tempering        bool   // If true, well19937c
	temperB, temperC uint32
	w, r, p          uint32
	maskU, maskL     uint32
	m1, m2, m3       uint32
	z0, z1, z2       uint32
}

func NewWELL19937a(seeds [624]uint32) (r well19937) {
	r = well19937{
		state_i:   0,
		_case:     1,
		tempering: false,
		temperB:   0xE46E1700,
		temperC:   0x9B868000,
		w:         32,
		r:         624,
		p:         31,
		maskU:     0x7FFFFFFF, // (0xffffffff>>(w-p))
		maskL:     0x80000000, // {bitwise NOT}(maskU)
		m1:        70,
		m2:        179,
		m3:        449,
	}
	copy(r.state[:], seeds[:])
	return
}

func NewWELL19937c(seeds [624]uint32) (r well19937) {
	r = well19937{
		state_i:   0,
		_case:     1,
		tempering: true,
		temperB:   0xE46E1700,
		temperC:   0x9B868000,
		w:         32,
		r:         624,
		p:         31,
		maskU:     0x7FFFFFFF, // (0xffffffff>>(w-p))
		maskL:     0x80000000, // {bitwise NOT}(maskU)
		m1:        70,
		m2:        179,
		m3:        449,
	}
	copy(r.state[:], seeds[:])
	return
}
