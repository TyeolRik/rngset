package rngset

import (
	"fmt"
	"log"
	"os"
)

type dieharder struct {
	rngName        string
	theNumberOfRNG uint64

	initialSeed uint64
}

func NewDieHarder(rngName string, theNumberOfRNG uint64, initialSeed uint64) dieharder {
	return dieharder{
		rngName:        rngName,
		theNumberOfRNG: theNumberOfRNG,
	}
}

func (d *dieharder) MakeFile(outputPath string) {
	fd, err := os.Create(outputPath + d.rngName + "_" + fmt.Sprintf("%v.txt", d.theNumberOfRNG))
	if err != nil {
		log.Fatalln("Failed to os.Create " + d.rngName + "_" + fmt.Sprintf("%v", d.theNumberOfRNG))
	}
	defer fd.Close()
	fd.WriteString("#==================================================================\n")
	fd.WriteString("# generator " + d.rngName + "  seed = " + fmt.Sprintf("%v\n", d.initialSeed))
	fd.WriteString("#==================================================================\n")
	fd.WriteString("type: d\n")
	fd.WriteString(fmt.Sprintf("count: %v\n", d.theNumberOfRNG))
	fd.WriteString("numbit: 64\n")

	test := NewMT19937_64(0)

	var i uint64 = 0
	for i = 0; i < d.theNumberOfRNG; i++ {
		fd.WriteString(fmt.Sprintf("%v\n", test.NextUint64()))
	}
}
