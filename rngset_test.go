package rngset_test

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/tyeolrik/rngset"
)

func TestSomething(t *testing.T) {
	lowerMask := uint32((1 << 31) - 1)
	upperMask := ^lowerMask
	fmt.Printf("lowerMask : 0x%08x\n", lowerMask)
	fmt.Printf("upperMask : 0x%08x\n", upperMask)
}

func TestDieharderOutput(t *testing.T) {
	d := rngset.NewDieHarder("sr__mt19937_64__well19937a_6block", 100000000, 0) // Why 100000000? https://webspace.science.uu.nl/~sleij101/Opgaven/LabClass/site/asm_diehard.php
	// Why 100000000 ? // 1억개 data
	// Not to rewound data. Dieharder test suite reuse previous data if its amount is small.
	d.MakeFile("./")
}

// 18시 20분 25초 시작
// dieharder -a -g 202 -f mt19937_64_100000000.txt && echo "test mail for sendmail gmail relay" | mail -s "Test End" kino6147@gmail.com && date
// date && dieharder -a -g 202 -f Well19937a_100000000.txt >> ./output/default_well19937a.txt && date

func TestDieharder6Block(t *testing.T) {
	d := rngset.NewDieHarder("sr__mt19937_64__well19937a", 10000000, 0)
	//d.MakeFileForBlockRand("./", 1, 32)
	//d.MakeFileForBlockRand("./", 2, 32)
	//d.MakeFileForBlockRand("./", 3, 32)
	//d.MakeFileForBlockRand("./", 4, 32)
	//d.MakeFileForBlockRand("./", 5, 32)
	d.MakeFileForBlockRand("./", 6, 32)
}

func TestDieharderKiss(t *testing.T) {
	// 1000만개 25초
	d := rngset.NewDieHarder("KISS", 100000000, 0)
	d.MakeFileForKISS("./generated/", 0, 64)
}

func TestDieharderWichmannHill(t *testing.T) {
	d := rngset.NewDieHarder("WichmannHill", 10000000, 0)
	d.MakeFileForWichmannHill("./", 0, 32)
}

func TestDieharderWell19937a(t *testing.T) {
	d := rngset.NewDieHarder("Well19937a", 100, 0)
	d.MakeFileForWell19937a("./")
}

func TestDieharderSR__WichmannHill__WichmannHill(t *testing.T) {
	d := rngset.NewDieHarder("sr__WichmannHill__WichmannHill", 100, 0)

	var i uint16
	for i = 1; i <= 6; i++ {
		d.MakeFileForSR__WichmannHill__WichmannHill("./", i, 64)
	}
}

func TestWELL512(t *testing.T) {
	b := make([]byte, 4)
	var seeds [16]uint32
	for i := 0; i < 16; i++ {
		rand.Read(b)
		seeds[i] = binary.LittleEndian.Uint32(b)
	}
	r := rngset.NewWELL512a(seeds)
	for i := 0; i < 10; i++ {
		fmt.Println(r.NewUint32())
	}
}

func TestDieharderSR__Kiss__WELL512(t *testing.T) {
	// go test -run TestDieharderSR__Kiss__WELL512 -timeout 10s
	d := rngset.NewDieHarder("sr__Kiss__WELL512", 5000000, 0)
	d.MakeFileForSR__Kiss__WELL512("./", 8, 32)
}

func TestDieharderSR__Kiss__WELL19937(t *testing.T) {
	// go test -run TestDieharderSR__Kiss__WELL19937 -timeout 500m
	// Test Failed // Too Slow
	d := rngset.NewDieHarder("sr__Kiss__WELL19937", 5000000, 0)
	d.MakeFileForSR__Kiss__WELL19937("./", 6, 32)
}
