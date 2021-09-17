package rngset

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"time"
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
	fd, err := os.Create(outputPath + d.rngName + "_" + fmt.Sprintf("%v.dat", d.theNumberOfRNG))
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

func (d *dieharder) MakeFileForWell19937a(outputPath string) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%v.dat", outputPath, d.rngName, d.theNumberOfRNG))
	if err != nil {
		log.Fatalln("Failed to os.Create " + d.rngName + "_" + fmt.Sprintf("%v", d.theNumberOfRNG))
	}
	defer fd.Close()
	fd.WriteString("#==================================================================\n")
	//	fd.WriteString(fmt.Sprintf(""))
	fd.WriteString(fmt.Sprintf("# generator %v  seed = %v\n", d.rngName, d.initialSeed))
	fd.WriteString("#==================================================================\n")
	fd.WriteString("type: d\n")
	fd.WriteString(fmt.Sprintf("count: %v\n", d.theNumberOfRNG))
	fd.WriteString(fmt.Sprintf("numbit: %v\n", 32))

	var seeds [624]uint32
	max_uint32 := ^uint32(0)
	for i := range seeds {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(max_uint32)))
		seeds[i] = uint32(r.Int64())
	}
	test := NewWELL19937a(seeds)

	var writeCount uint64 = 0
	for writeCount < d.theNumberOfRNG {
		fd.WriteString(fmt.Sprintf("%v\n", test.NewUint32()))
		writeCount++
	}
}

func (d *dieharder) MakeFileForBlockRand(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
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
		test := NewSR__mt19937_64__well19937a(blockSize)
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

func (d *dieharder) MakeFileForKISS(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
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
	s1 := binary.LittleEndian.Uint64(b)
	rand.Read(b)
	s2 := binary.LittleEndian.Uint64(b)
	test := NewKISS(s1, s2)

	var writeCount uint64 = 0
	w := bufio.NewWriter(fd)
	for writeCount < d.theNumberOfRNG {
		fmt.Fprintf(w, "%v\n", test.NextUInt64())
		w.Flush()
		writeCount++
	}
}

func (d *dieharder) MakeFileForWichmannHill(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
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

func (d *dieharder) MakeFileForSR__WichmannHill__WichmannHill(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
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
		test := NewSR__WichmannHill__WichmannHill(blockSize)
		useridIndex := 0
		var isAllMined bool = false

		var blockIndex uint16
		for blockIndex = 0; blockIndex < blockSize; blockIndex++ {
			for i := 0; i < 3; i++ {
				tempRandom, err := rand.Int(rand.Reader, big.NewInt(int64(^uint32(0)>>1)))
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

func (d *dieharder) MakeFileForSR__Kiss__WELL512(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
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
	b := make([]byte, 8)
	w := bufio.NewWriter(fd)

	for writeCount < d.theNumberOfRNG {
		test := NewSR__Kiss__WELL512(blockSize)
		useridIndex := 0
		var isAllMined bool = false

		var blockIndex uint16
		for blockIndex = 0; blockIndex < blockSize; blockIndex++ {
			for i := 0; i < 2; i++ {
				rand.Read(b)
				tempRandom := binary.LittleEndian.Uint64(b)
				test.Participate(strconv.Itoa(useridIndex), tempRandom)
				useridIndex++
			}
			isAllMined = test.Mining()
		}

		for !isAllMined {
			isAllMined = test.Mining()
		}

		getData := test.GetReturns(16)

		for i := 0; i < len(getData); i++ {
			fmt.Fprintf(w, "%v\n", getData[i])
			w.Flush()
			// fd.WriteString(fmt.Sprintf("%v\n", getData[i]))
			writeCount++
		}
	}
}

func (d *dieharder) MakeFileForSR__Kiss__WELL512__Extreme_8block_2participant(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
	w := bufio.NewWriter(fd)
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
	b := make([]byte, 8)
	participantInput := make([]uint64, 16)
	var wellSeed [16]uint32

	var for1, for2 int64

	for writeCount < d.theNumberOfRNG {
		for i := range participantInput {
			rand.Read(b)
			participantInput[i] = binary.LittleEndian.Uint64(b)
		}
		check1 := time.Now().UnixNano()
		for i := 0; i < 16; i = i + 2 {
			_newKISS := NewKISS(participantInput[i], participantInput[i+1])
			temp := _newKISS.NextUInt64()
			wellSeed[i] = uint32(temp >> 32)
			wellSeed[i+1] = uint32((temp << 32) >> 32)
		}
		check2 := time.Now().UnixNano()
		_newWELL512 := NewWELL512a(wellSeed)
		for i := 0; i < 16; i++ {
			fmt.Fprintf(w, "%v\n", _newWELL512.NewUint32())
		}
		check3 := time.Now().UnixNano()
		for1 := for1 + (check2 - check1)
		for2 := for2 + (check3 - check2)
		fmt.Println("For1 :", for1)
		fmt.Println("For2 :", for2)
	}

}

// This takes too many time
func (d *dieharder) MakeFileForSR__Kiss__WELL19937(outputPath string, blockSize uint16, bits_32_or_64 uint16) {
	fd, err := os.Create(fmt.Sprintf("%v%v_%vBlock_%v.dat", outputPath, d.rngName, blockSize, d.theNumberOfRNG))
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
	b := make([]byte, 8)

	for writeCount < d.theNumberOfRNG {
		test := NewSR__Kiss__WELL19937(blockSize)
		useridIndex := 0
		var isAllMined bool = false

		var blockIndex uint16
		for blockIndex = 0; blockIndex < blockSize; blockIndex++ {
			for i := 0; i < 4; i++ {
				rand.Read(b)
				tempRandom := binary.LittleEndian.Uint64(b)
				test.Participate(strconv.Itoa(useridIndex), tempRandom)
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
