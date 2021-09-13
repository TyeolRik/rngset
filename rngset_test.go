package rngset_test

import (
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
	d := rngset.NewDieHarder("mt19937_64", 100000000, 0) // Why 100000000? https://webspace.science.uu.nl/~sleij101/Opgaven/LabClass/site/asm_diehard.php
	// Why 100000000 ? // 1억개 data
	// Not to rewound data. Dieharder test suite reuse previous data if its amount is small.
	d.MakeFile("./")
}

// 18시 20분 25초 시작
// dieharder -a -g 202 -f mt19937_64_100000000.txt && echo "test mail for sendmail gmail relay" | mail -s "Test End" kino6147@gmail.com && date

func TestDieharderWichMannHill(t *testing.T) {
	d := rngset.NewDieHarder("WichMann_Hill", 10000000, 0)
	
}