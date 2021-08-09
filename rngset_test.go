package rngset_test

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	lowerMask := uint64((1 << 31) - 1)
	upperMask := ^lowerMask
	fmt.Printf("lowerMask : 0x%016x\n", lowerMask)
	fmt.Printf("upperMask : 0x%016x\n", upperMask)
}
