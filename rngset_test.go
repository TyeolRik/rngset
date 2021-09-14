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

func TestDieharder6Block(t *testing.T) {
	d := rngset.NewDieHarder("sr__mt19937_64__well19937a", 10000000, 0)
	//d.MakeFileForBlockRand("./", 1, 32)
	//d.MakeFileForBlockRand("./", 2, 32)
	//d.MakeFileForBlockRand("./", 3, 32)
	//d.MakeFileForBlockRand("./", 4, 32)
	//d.MakeFileForBlockRand("./", 5, 32)
	d.MakeFileForBlockRand("./", 6, 32)
}

func TestDieharderWichmannHill(t *testing.T) {
	d := rngset.NewDieHarder("WichmannHill", 10000000, 0)
	d.MakeFileForWichmannHill("./", 0, 32)
}

func TestCryptoRand(t *testing.T) {
	b := make([]byte, 8)
	rand.Read(b)
	_s1 := binary.LittleEndian.Uint64(b)
	rand.Read(b)
	_s2 := binary.LittleEndian.Uint64(b)
	rand.Read(b)
	_s3 := binary.LittleEndian.Uint64(b)
	fmt.Println(_s1)
	fmt.Println(_s2)
	fmt.Println(_s3)
}
