package rngset

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	fmt.Printf("%b\n", 7)
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
