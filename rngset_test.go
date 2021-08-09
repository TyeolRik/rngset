package rngset_test

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	lowerMask := uint32((1 << 31) - 1)
	upperMask := ^lowerMask
	fmt.Printf("lowerMask : 0x%08x\n", lowerMask)
	fmt.Printf("upperMask : 0x%08x\n", upperMask)
}

func TestBit(t *testing.T) {
	var maskU uint32 = 0xffffffff >> 1
	var maskL uint32 = ^maskU
	fmt.Printf("maskU: 0x%08X\n", maskU)
	fmt.Printf("maskL: 0x%08X\n", maskL)
}
