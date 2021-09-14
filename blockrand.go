package rngset

import (
	"math/rand"
)

type participant struct {
	userID    string
	userInput uint64

	returns uint32
}

type block struct {
	realSeed uint64

	participants []participant
}

type sr__mt19937_64__well19937a struct {
	// Information
	_blockSize uint16 // n개의 블록이 모이면 output을 만들 수 있다.

	state  uint64  // 현재 단계
	blocks []block // n개의 블록이 모이면 output을 만들 수 있음.

	participantCounter uint32
}

func NewSR(blockSize uint16) (ret sr__mt19937_64__well19937a) {
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
		var allReturns []uint32 = make([]uint32, r.participantCounter)
		for i := range allReturns {
			allReturns[i] = uint32(well.NextFloat64() * float64(^uint32(0)))
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

func (r *sr__mt19937_64__well19937a) GetFirst3Returns() (ret [3]uint32) {
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
