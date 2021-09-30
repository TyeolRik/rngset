package rngset

import (
	"encoding/binary"
	"math/rand"

	"github.com/ethereum/go-ethereum/crypto"
)

type participant struct {
	userID    string
	userInput uint64

	returns uint64
}

type block struct {
	realSeed uint64

	participants []participant
}

type keccakBlock struct {
	realSeed     []byte
	participants []participant
}

type sr__mt19937_64__well19937a struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64  // 현재 단계
	blocks []block // n개의 블록이 모이면 output을 만들 수 있음.

	participantCounter uint32
}

func NewSR__mt19937_64__well19937a(blockSize uint16) (ret sr__mt19937_64__well19937a) {
	ret = sr__mt19937_64__well19937a{
		_blockSize: blockSize,
		state:      0,
	}
	ret.blocks = make([]block, blockSize)
	for i := range ret.blocks {
		ret.blocks[i].realSeed = 0
		ret.blocks[i].participants = make([]participant, 0)
	}
	return
}

func (r *sr__mt19937_64__well19937a) Participate(userID string, userInput uint64) {
	r.blocks[r.state].participants = append(r.blocks[r.state].participants, participant{userID: userID, userInput: userInput})
	r.participantCounter++
}

func (r *sr__mt19937_64__well19937a) Mining() bool {
	if r.state == uint64(r._blockSize) {
		// Make Real Output (Real Random Number)
		seedPadding := [624]uint32{}
		for i := range seedPadding {
			seedPadding[i] = uint32(r.blocks[r.state%uint64(r._blockSize)].realSeed) // Block 1,2,3,1,2,3,1,2,3 이런 식으로 반복적으로 채움
		}
		well := NewWELL19937a(seedPadding)
		// Since 0 <= output < 1, 0 * 2^32-1 <= output * 2^32-1 < 1 * 2^32-1
		var allReturns []uint64 = make([]uint64, r.participantCounter)
		for i := range allReturns {
			allReturns[i] = uint64(well.NextFloat64() * float64(^uint64(0)))
		}
		rand.Seed(int64(well.NextFloat64() * float64(^uint64(0))))
		rand.Shuffle(len(allReturns), func(i int, j int) { allReturns[i], allReturns[j] = allReturns[j], allReturns[i] })
		globalIndex := 0
		for _, eachBlock := range r.blocks {
			for i := range eachBlock.participants {
				eachBlock.participants[i].returns = allReturns[globalIndex]
				globalIndex++
			}
		}
		return true
	} else {
		// Calculate realSeed
		// MT19937_64 needs only 1 seed. So, all participant input should be XORed. (For resulting only ONE value)
		for i := range r.blocks[r.state].participants {
			r.blocks[r.state].realSeed = r.blocks[r.state].realSeed ^ r.blocks[r.state].participants[i].userInput
		}
		mt19937_64 := NewMT19937_64(r.blocks[r.state].realSeed)
		r.blocks[r.state].realSeed = mt19937_64.NextUint64()

		// Go Next Block State
		r.state++

		return false
	}
}

func (r *sr__mt19937_64__well19937a) GetFirst3Returns() (ret [3]uint64) {
	globalIndex := 0
	for _, eachBlock := range r.blocks {
		for i := range eachBlock.participants {
			ret[globalIndex] = eachBlock.participants[i].returns
			globalIndex++

			if globalIndex >= 3 {
				return
			}
		}
	}
	return
}

type sr__WichmannHill__WichmannHill struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64  // 현재 단계
	blocks []block // n개의 블록이 모이면 output을 만들 수 있음.

	participantCounter uint32
}

func NewSR__WichmannHill__WichmannHill(blockSize uint16) (ret sr__WichmannHill__WichmannHill) {
	blocks := make([]block, blockSize)
	for i := range blocks {
		blocks[i].realSeed = 0
		blocks[i].participants = make([]participant, 0)
	}
	ret = sr__WichmannHill__WichmannHill{
		_blockSize:         blockSize,
		state:              0,
		blocks:             blocks,
		participantCounter: 0,
	}
	return
}

func (r *sr__WichmannHill__WichmannHill) Participate(userID string, userInput uint64) {
	r.blocks[r.state].participants = append(r.blocks[r.state].participants, participant{userID: userID, userInput: userInput})
	r.participantCounter++
}

func (r *sr__WichmannHill__WichmannHill) Mining() bool {
	if r.state == uint64(r._blockSize) {
		// Make Real Output (Real Random Number)
		// 각 블록에 대한 realSeed를 만든다.
		for blockIndex, block := range r.blocks {
			var blockseed [3]uint64
			if len(block.participants) < 3 {
				// 각 Block에 참여자가 적으면 MT19937_64로 늘려준다.
				switch len(block.participants) {
				case 1:
					blockseed[0] = block.participants[0].userInput
					temp := NewMT19937_64(blockseed[0])
					blockseed[1] = temp.NextUint64()
					blockseed[2] = temp.NextUint64()
				case 2:
					blockseed[0] = block.participants[0].userInput
					blockseed[1] = block.participants[1].userInput
					temp := NewMT19937_64(blockseed[0] ^ blockseed[1])
					blockseed[2] = temp.NextUint64()
				}
			} else if len(block.participants) == 3 {
				blockseed[0] = block.participants[0].userInput
				blockseed[1] = block.participants[1].userInput
				blockseed[2] = block.participants[2].userInput
			} else {
				// 각 Block에 참여자가 많으면 모듈러로 XOR 시킨다.
				for part_index, participant := range block.participants {
					blockseed[part_index%3] = blockseed[part_index%3] ^ participant.userInput
				}
			}
			makeRealseed_byWichmannHill := NewWichmannHill(int64(blockseed[0]), int64(blockseed[1]), int64(blockseed[2]))
			// 0 <= target <= ^uint(0)
			makeRealseed_byWichmannHill.Generate()
			r.blocks[blockIndex].realSeed = uint64(float64(^uint64(0)) * makeRealseed_byWichmannHill.Generate())
		}
		var realSeed [3]uint64
		if r._blockSize == 1 {
			// 각 Block에 참여자가 적으면 MT19937_64로 늘려준다.
			realSeed[0] = r.blocks[0].realSeed
			temp := NewMT19937_64(realSeed[0])
			realSeed[1] = temp.NextUint64()
			realSeed[2] = temp.NextUint64()
		} else if r._blockSize == 2 {
			realSeed[0] = r.blocks[0].realSeed
			realSeed[1] = r.blocks[1].realSeed
			temp := NewMT19937_64(realSeed[0] ^ realSeed[1])
			realSeed[2] = temp.NextUint64()
		} else if r._blockSize == 3 {
			realSeed[0] = r.blocks[0].realSeed
			realSeed[1] = r.blocks[1].realSeed
			realSeed[2] = r.blocks[2].realSeed
		} else {
			// 각 Block에 참여자가 많으면 모듈러로 XOR 시킨다.
			for block_index, block := range r.blocks {
				realSeed[block_index%3] = realSeed[block_index%3] ^ block.realSeed
			}
		}
		wichmann := NewWichmannHill(int64(realSeed[0]), int64(realSeed[1]), int64(realSeed[2]))
		wichmann.Generate()

		var allReturns []uint64 = make([]uint64, r.participantCounter)
		for i := range allReturns {
			allReturns[i] = uint64(wichmann.Generate() * float64(^uint64(0)))
		}
		rand.Seed(int64(wichmann.Generate() * float64(^uint64(0))))
		rand.Shuffle(len(allReturns), func(i int, j int) { allReturns[i], allReturns[j] = allReturns[j], allReturns[i] })
		globalIndex := 0
		for _, eachBlock := range r.blocks {
			for i := range eachBlock.participants {
				eachBlock.participants[i].returns = allReturns[globalIndex]
				globalIndex++
			}
		}
		return true
	} else {
		// Go Next Block State
		r.state++

		return false
	}
}

func (r *sr__WichmannHill__WichmannHill) GetFirst3Returns() (ret [3]uint64) {
	globalIndex := 0
	for _, eachBlock := range r.blocks {
		for i := range eachBlock.participants {
			ret[globalIndex] = eachBlock.participants[i].returns
			globalIndex++

			if globalIndex >= 3 {
				return
			}
		}
	}
	return
}

type sr__Kiss__WELL512 struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64  // 현재 단계
	blocks []block // n개의 블록이 모이면 output을 만들 수 있음.

	participantCounter uint32
}

func NewSR__Kiss__WELL512(blockSize uint16) (ret sr__Kiss__WELL512) {
	blocks := make([]block, blockSize)
	for i := range blocks {
		blocks[i].realSeed = 0
		blocks[i].participants = make([]participant, 0)
	}
	ret = sr__Kiss__WELL512{
		_blockSize:         blockSize,
		state:              0,
		blocks:             blocks,
		participantCounter: 0,
	}
	return
}

func (r *sr__Kiss__WELL512) Participate(userID string, userInput uint64) {
	r.blocks[r.state].participants = append(r.blocks[r.state].participants, participant{userID: userID, userInput: userInput})
	r.participantCounter++
}

func (r *sr__Kiss__WELL512) Mining() bool {
	var seedLength int = 2
	if r.state == uint64(r._blockSize) {
		// Make Real Output (Real Random Number)
		// 각 블록에 대한 realSeed를 만든다.
		for blockIndex, block := range r.blocks {
			var blockseed [2]uint64
			if len(block.participants) == 1 {
				// 각 Block에 참여자가 적으면 MT19937_64로 늘려준다.
				switch len(block.participants) {
				case 1:
					blockseed[0] = block.participants[0].userInput
					temp := NewMT19937_64(blockseed[0])
					blockseed[1] = temp.NextUint64()
				}
			} else if len(block.participants) == 2 {
				blockseed[0] = block.participants[0].userInput
				blockseed[1] = block.participants[1].userInput
			} else {
				// 각 Block에 참여자가 많으면 모듈러로 XOR 시킨다.
				for part_index, participant := range block.participants {
					blockseed[part_index%seedLength] = blockseed[part_index%seedLength] ^ participant.userInput
				}
			}
			makeRealseed_byKISS := NewKISS(blockseed[0], blockseed[1])
			// 0 <= target <= ^uint(0)
			makeRealseed_byKISS.NextUInt64()
			r.blocks[blockIndex].realSeed = makeRealseed_byKISS.NextUInt64()
		}

		// Using WELL19937
		var realSeed [16]uint32

		for blockIndex := range r.blocks {
			realSeed[2*blockIndex] = uint32(r.blocks[blockIndex].realSeed >> 32)
			realSeed[2*blockIndex+1] = uint32(r.blocks[blockIndex].realSeed & 0x00000000FFFFFFFF)
		}
		for realSeedIndex := 2 * len(r.blocks); realSeedIndex < len(realSeed); realSeedIndex = realSeedIndex + 2 {
			temp := NewKISS(uint64(realSeed[realSeedIndex-2]), uint64(realSeed[realSeedIndex-1]))
			newUint64 := temp.NextUInt64()
			realSeed[realSeedIndex] = uint32(newUint64 >> 32)
			realSeed[realSeedIndex+1] = uint32(newUint64 & 0x00000000FFFFFFFF)
		}

		/*
			for blockIndex := range r.blocks {
				realSeed[2*blockIndex] = uint32(r.blocks[blockIndex].realSeed >> 32)
				realSeed[2*blockIndex+1] = uint32(r.blocks[blockIndex].realSeed & 0x00000000FFFFFFFF)
			}
			for realSeedIndex := len(r.blocks); realSeedIndex < len(realSeed); realSeedIndex++ {
				realSeed[realSeedIndex] = realSeed[realSeedIndex%len(r.blocks)]
			}
		*/
		well512 := NewWELL512a(realSeed)

		var allReturns []uint64 = make([]uint64, r.participantCounter)
		for i := range allReturns {
			allReturns[i] = uint64(well512.NewUint32())
		}
		rand.Seed(int64(well512.NewUint32()))
		rand.Shuffle(len(allReturns), func(i int, j int) { allReturns[i], allReturns[j] = allReturns[j], allReturns[i] })
		globalIndex := 0
		for _, eachBlock := range r.blocks {
			for i := range eachBlock.participants {
				eachBlock.participants[i].returns = allReturns[globalIndex]
				globalIndex++
			}
		}
		return true
	} else {
		// Go Next Block State
		r.state++

		return false
	}
}

func (r *sr__Kiss__WELL512) GetReturns(theNumberOfReturns int) (ret []uint64) {
	globalIndex := 0
	ret = make([]uint64, theNumberOfReturns)
	for _, eachBlock := range r.blocks {
		for i := range eachBlock.participants {
			ret[globalIndex] = eachBlock.participants[i].returns
			globalIndex++

			if globalIndex >= len(ret) {
				return
			}
		}
	}
	return
}

func (r *sr__Kiss__WELL512) GetFirst18Returns() (ret [18]uint64) {
	globalIndex := 0
	for _, eachBlock := range r.blocks {
		for i := range eachBlock.participants {
			ret[globalIndex] = eachBlock.participants[i].returns
			globalIndex++

			if globalIndex >= len(ret) {
				return
			}
		}
	}
	return
}

type sr__Kiss__WELL19937 struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64  // 현재 단계
	blocks []block // n개의 블록이 모이면 output을 만들 수 있음.

	participantCounter uint32
}

func NewSR__Kiss__WELL19937(blockSize uint16) (ret sr__Kiss__WELL19937) {
	blocks := make([]block, blockSize)
	for i := range blocks {
		blocks[i].realSeed = 0
		blocks[i].participants = make([]participant, 0)
	}
	ret = sr__Kiss__WELL19937{
		_blockSize:         blockSize,
		state:              0,
		blocks:             blocks,
		participantCounter: 0,
	}
	return
}

func (r *sr__Kiss__WELL19937) Participate(userID string, userInput uint64) {
	r.blocks[r.state].participants = append(r.blocks[r.state].participants, participant{userID: userID, userInput: userInput})
	r.participantCounter++
}

func (r *sr__Kiss__WELL19937) Mining() bool {
	var seedLength int = 2
	if r.state == uint64(r._blockSize) {
		// Make Real Output (Real Random Number)
		// 각 블록에 대한 realSeed를 만든다.
		for blockIndex, block := range r.blocks {
			var blockseed [2]uint64
			if len(block.participants) == 1 {
				// 각 Block에 참여자가 적으면 MT19937_64로 늘려준다.
				switch len(block.participants) {
				case 1:
					blockseed[0] = block.participants[0].userInput
					temp := NewMT19937_64(blockseed[0])
					blockseed[1] = temp.NextUint64()
				}
			} else if len(block.participants) == 2 {
				blockseed[0] = block.participants[0].userInput
				blockseed[1] = block.participants[1].userInput
			} else {
				// 각 Block에 참여자가 많으면 모듈러로 XOR 시킨다.
				for part_index, participant := range block.participants {
					blockseed[part_index%seedLength] = blockseed[part_index%seedLength] ^ participant.userInput
				}
			}
			makeRealseed_byKISS := NewKISS(blockseed[0], blockseed[1])
			// 0 <= target <= ^uint(0)
			makeRealseed_byKISS.NextUInt64()
			r.blocks[blockIndex].realSeed = makeRealseed_byKISS.NextUInt64()
		}

		// Using WELL19937
		var realSeed [624]uint32
		for blockIndex := range r.blocks {
			realSeed[2*blockIndex] = uint32(r.blocks[blockIndex].realSeed >> 32)
			realSeed[2*blockIndex+1] = uint32(r.blocks[blockIndex].realSeed & 0x00000000FFFFFFFF)
		}
		for realSeedIndex := 2 * len(r.blocks); realSeedIndex < 624; realSeedIndex = realSeedIndex + 2 {
			// Part Meaning
			// 1234 5678 9ABC DEF0
			//  p1   p2   p3   p4
			// 	 block1    block2
			//   rseed     rseed
			part1 := realSeed[realSeedIndex-2] & 0xFFFF0000
			part2 := realSeed[realSeedIndex-2] & 0x0000FFFF
			part3 := realSeed[realSeedIndex-1] & 0xFFFF0000
			part4 := realSeed[realSeedIndex-1] & 0x0000FFFF
			realSeed[realSeedIndex] = part1 | part4
			realSeed[realSeedIndex+1] = part2 | part3
		}
		well19937a := NewWELL19937a(realSeed)

		var allReturns []uint64 = make([]uint64, r.participantCounter)
		for i := range allReturns {
			allReturns[i] = uint64(well19937a.NewUint32())
		}
		rand.Seed(int64(well19937a.NewUint32()))
		rand.Shuffle(len(allReturns), func(i int, j int) { allReturns[i], allReturns[j] = allReturns[j], allReturns[i] })
		globalIndex := 0
		for _, eachBlock := range r.blocks {
			for i := range eachBlock.participants {
				eachBlock.participants[i].returns = allReturns[globalIndex]
				globalIndex++
			}
		}
		return true
	} else {
		// Go Next Block State
		r.state++

		return false
	}
}

func (r *sr__Kiss__WELL19937) GetFirst3Returns() (ret [3]uint64) {
	globalIndex := 0
	for _, eachBlock := range r.blocks {
		for i := range eachBlock.participants {
			ret[globalIndex] = eachBlock.participants[i].returns
			globalIndex++

			if globalIndex >= 3 {
				return
			}
		}
	}
	return
}

type sr__Keccak256__WELL512a struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64        // 현재 단계
	blocks []keccakBlock // n개의 블록이 모이면 output을 만들 수 있음. realSeed = 256bits

	participantCounter uint32
}

func NewSR__Keccak256__WELL512a(blockSize uint16) (ret sr__Keccak256__WELL512a) {
	blocks := make([]keccakBlock, blockSize)
	for i := range blocks {
		blocks[i].participants = make([]participant, 0)
	}
	ret = sr__Keccak256__WELL512a{
		_blockSize:         blockSize,
		state:              0,
		blocks:             blocks,
		participantCounter: 0,
	}
	return
}

func (r *sr__Keccak256__WELL512a) Participate(userID string, userInput uint64) {
	r.blocks[r.state].participants = append(r.blocks[r.state].participants, participant{userID: userID, userInput: userInput})
	r.participantCounter++
}

func (r *sr__Keccak256__WELL512a) Mining() bool {
	tempByte := make([]byte, 8)
	if r.state == uint64(r._blockSize) {
		// Make Real Output (Real Random Number)
		// 각 블록에 대한 realSeed를 만든다.
		for blockIndex, block := range r.blocks {
			var blockseed uint64 = 0
			for _, participant := range block.participants {
				blockseed = blockseed ^ participant.userInput
			}
			binary.LittleEndian.PutUint64(tempByte, blockseed)
			hashValue := crypto.Keccak256(tempByte)
			r.blocks[blockIndex].realSeed = hashValue
		}

		// Using WELL512a
		var realSeed [16]uint32
		for blockIndex, eachBlock := range r.blocks {
			for i := 0; i < 8; i++ {
				realSeed[8*(blockIndex&1)+i] = binary.LittleEndian.Uint32(eachBlock.realSeed[4*i : 4*(i+1)])
			}
		}
		if len(r.blocks) == 1 {
			anotherHash := crypto.Keccak256(r.blocks[0].realSeed)
			for i := 0; i < 8; i++ {
				realSeed[8+i] = binary.LittleEndian.Uint32(anotherHash[4*i : 4*(i+1)])
			}
		}
		well512a := NewWELL512a(realSeed)

		var allReturns []uint64 = make([]uint64, r.participantCounter)
		for i := range allReturns {
			allReturns[i] = uint64(well512a.NewUint32())
		}
		rand.Seed(int64(well512a.NewUint32()))
		rand.Shuffle(len(allReturns), func(i int, j int) { allReturns[i], allReturns[j] = allReturns[j], allReturns[i] })
		globalIndex := 0
		for _, eachBlock := range r.blocks {
			for i := range eachBlock.participants {
				eachBlock.participants[i].returns = allReturns[globalIndex]
				globalIndex++
			}
		}
		return true
	} else {
		// Go Next Block State
		r.state++

		return false
	}
}

func (r *sr__Keccak256__WELL512a) GetReturns(theNumberOfReturns int) (ret []uint32) {
	globalIndex := 0
	ret = make([]uint32, theNumberOfReturns)
	for _, eachBlock := range r.blocks {
		for i := range eachBlock.participants {
			ret[globalIndex] = uint32(eachBlock.participants[i].returns)
			globalIndex++

			if globalIndex >= len(ret) {
				return
			}
		}
	}
	return
}
