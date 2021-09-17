package rngset

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestBinary(t *testing.T) {
	a := ^uint64(0) - 58
	fmt.Println(a)
	fmt.Println(" Max:", ^uint64(0))
	fmt.Println("Diff:", ^uint64(0)-a)
	a = a + 59 // Overflow = 0 // Masking Naturally.
	fmt.Println("Test:", a)
}

func TestMiddleSquare(t *testing.T) {
	r := NewMiddleSquare(42)
	for {
		now := r.Generate()
		if now == 0 {
			break
		}
		fmt.Println(now)
	}

}

func BenchmarkKISS(b *testing.B) {
	randB := make([]byte, 8)
	var r1, r2 uint64
	rand.Read(randB)
	r1 = binary.LittleEndian.Uint64(randB)
	rand.Read(randB)
	r2 = binary.LittleEndian.Uint64(randB)
	for i := 0; i < b.N; i++ {
		_kiss := NewKISS(r1, r2)
		_kiss.NextUInt64()
	}
}

func BenchmarkWELL512a(b *testing.B) {
	randB := make([]byte, 4)
	var seeds [16]uint32
	for i := range seeds {
		rand.Read(randB)
		seeds[i] = binary.LittleEndian.Uint32(randB)
	}
	for i := 0; i < b.N; i++ {
		well512a := NewWELL512a(seeds)
		well512a.NewUint32()
	}
}

func checkhash(t *testing.T, name string, f func([]byte) []byte, msg, exp []byte) {
	sum := f(msg)
	if !bytes.Equal(exp, sum) {
		t.Fatalf("hash %s mismatch: want: %x have: %x", name, exp, sum)
	}
}

func TestKeccak256(t *testing.T) {
	msg := []byte("12345")
	hash := crypto.Keccak256Hash(msg)
	fmt.Println(hash.Bytes())
	fmt.Println(len(hash.Bytes()))

	var temp []byte = []byte("1")
	fmt.Println(temp)
}

func BenchmarkKeccak256(b *testing.B) {
	randB := make([]byte, 8)
	rand.Read(randB)
	for i := 0; i < b.N; i++ {
		crypto.Keccak256Hash(randB)
	}
}
func BenchmarkKeccak512(b *testing.B) {
	randB := make([]byte, 8)
	rand.Read(randB)
	for i := 0; i < b.N; i++ {
		crypto.Keccak512(randB)
	}
}

func TestWichmannHill(t *testing.T) {
	r := NewWichmannHill(23415, 4903, 25333)
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	// fmt.Println(WichmannHill(23415, 4903, 25333))
}

func TestRule30(t *testing.T) {
	r := NewRule30(8, 11)
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
}

func TestICG(t *testing.T) {
	r := NewInversiveCongruentialGenerator(5, 2, 3, 1)
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
	fmt.Println(r.Generate())
}

func TestACORN(t *testing.T) {
	r := NewACORN(10)
	fmt.Println(r.NextFloat64())
	fmt.Println(r.NextFloat64())
	fmt.Println(r.NextFloat64())
	fmt.Println(r.NextFloat64())
	fmt.Println(r.NextFloat64())
	fmt.Println(r.NextFloat64())
}

func TestAWCandSWB(t *testing.T) {
	AWC := NewAWC_Recommend()
	fmt.Println("AWC")
	for i := 0; i < 5; i++ {
		fmt.Println(AWC.NextFloat64())
	}
	AWC_C := NewAWC_Recommend()
	fmt.Println("AWC_C")
	for i := 0; i < 5; i++ {
		fmt.Println(AWC_C.NextFloat64())
	}
	SWB1 := NewAWC_Recommend()
	fmt.Println("SWB1")
	for i := 0; i < 5; i++ {
		fmt.Println(SWB1.NextFloat64())
	}
	SWB2 := NewAWC_Recommend()
	fmt.Println("SWB2")
	for i := 0; i < 5; i++ {
		fmt.Println(SWB2.NextFloat64())
	}
}

func TestKISS(t *testing.T) {
	r := NewKISS(123456789987654321, 362436069362436069)
	var index uint64 = 0
	for {
		temp := r.NewFloat64()
		if temp > 0.9999999 {
			fmt.Println(index, ":", temp)
			break
		}
	}
	/*
		var temp uint64
		for i := 0; i < 1000000000; i++ {
			temp = r.NextUInt64()
		}
		fmt.Println("Result: ", temp)
		fmt.Println("Answer: ", 5033346742750153761)
	*/
}

func TestMT19937_64(t *testing.T) {
	r := NewMT19937_64(0)
	fmt.Println(r.NextUint64())
}

func TestWELL512a(t *testing.T) {
	r := NewWELL512a([16]uint32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	fmt.Println(r.NextFloat64())
}

func TestWELL19937a(t *testing.T) {
	seed := [624]uint32{}
	for i := range seed {
		seed[i] = uint32(i)
	}
	r := NewWELL19937a(seed)
	for i := 0; i < 20; i++ {
		fmt.Println(i, r.NextFloat64())
	}
}
