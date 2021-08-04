package rngset

import (
	"crypto/rand"
	"math/big"
)

const FLOAT_PRECISION = 1000000000 // Maximum of int32 = 2,147,483,647

func getCryptoRandInt(min, max int64) int64 {
	nBig, _ := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	n := nBig.Int64()
	return int64(n) + min
}

func getCryptoRandFloat(min, max float64) float64 {
	minInt := int64(min * FLOAT_PRECISION)
	maxInt := int64(max * FLOAT_PRECISION)

	return float64(getCryptoRandInt(minInt, maxInt)) / FLOAT_PRECISION
}
