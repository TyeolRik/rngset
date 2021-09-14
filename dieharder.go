package rngset

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
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

func (d *dieharder) MakeFileForBlockRand(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.txt", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
	if err != nil {
		log.Fatalln("Failed to os.Create " + d.rngName + "_" + fmt.Sprintf("%v", d.theNumberOfRNG))
	}
	defer fd.Close()
	fd.WriteString("#==================================================================\n")
	//	fd.WriteString(fmt.Sprintf(""))
	fd.WriteString(fmt.Sprintf("# generator %v_%vBlock  seed = %v\n", d.rngName, blockSize, d.initialSeed))
	fd.WriteString("#==================================================================\n")
	fd.WriteString("type: d\n")
	fd.WriteString(fmt.Sprintf("count: %v\n", d.theNumberOfRNG))
	fd.WriteString(fmt.Sprintf("numbit: %v\n", bits_32_or_64))

	var writeCount uint64 = 0

	for writeCount < d.theNumberOfRNG {
		test := NewSR(blockSize)
		useridIndex := 0
		var isAllMined bool = false

		// Input Real Random Golang CryptoRand = /dev/urandom
		var block uint16
		for block = 0; block < blockSize; block++ {
			for i := 0; i < 3; i++ {
				tempRandom, err := rand.Int(rand.Reader, big.NewInt(int64(^uint64(0)>>1)))
				if err != nil {
					log.Fatalln(err)
				}
				test.Participate(strconv.Itoa(useridIndex), tempRandom.Uint64())
				useridIndex++
			}
			isAllMined = test.Mining()
		}
		for !isAllMined {
			isAllMined = test.Mining()
		}

		first3data := test.GetFirst3Returns()

		for i := 0; i < 3; i++ {
			fd.WriteString(fmt.Sprintf("%v\n", first3data[i]))
			writeCount++
			if writeCount >= d.theNumberOfRNG {
				break
			}
		}
	}

}

func (d *dieharder) MakeFileForWichmannHill(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.txt", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
	if err != nil {
		log.Fatalln("Failed to os.Create " + d.rngName + "_" + fmt.Sprintf("%v", d.theNumberOfRNG))
	}
	defer fd.Close()
	fd.WriteString("#==================================================================\n")
	//	fd.WriteString(fmt.Sprintf(""))
	fd.WriteString(fmt.Sprintf("# generator %v_%vBlock  seed = %v\n", d.rngName, blockSize, d.initialSeed))
	fd.WriteString("#==================================================================\n")
	fd.WriteString("type: d\n")
	fd.WriteString(fmt.Sprintf("count: %v\n", d.theNumberOfRNG))
	fd.WriteString(fmt.Sprintf("numbit: %v\n", bits_32_or_64))

	b := make([]byte, 8)
	rand.Read(b)
	_s1 := int64(binary.LittleEndian.Uint64(b))
	rand.Read(b)
	_s2 := int64(binary.LittleEndian.Uint64(b))
	rand.Read(b)
	_s3 := int64(binary.LittleEndian.Uint64(b))

	test := NewWichmannHill(_s1, _s2, _s3)
	// 0 <= return <= 2^32-1
	var MAX_uint32 uint32 = ^uint32(0)

	var writeCount uint64 = 0
	for writeCount < d.theNumberOfRNG {
		generated := make([]uint32, 18)
		for i := range generated {
			generated[i] = uint32(test.Generate() * float64(MAX_uint32))
		}
		// Fisher-Yates Shuffle
		for i := len(generated) - 1; i > 0; i-- {
			j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
			generated[i], generated[j.Int64()] = generated[j.Int64()], generated[i]
		}
		for i := 0; i < 3; i++ {
			fd.WriteString(fmt.Sprintf("%v\n", generated[i]))
			writeCount++
			if writeCount >= d.theNumberOfRNG {
				break
			}
		}
	}
}
