package rngset

import (
	"fmt"
	"testing"
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
