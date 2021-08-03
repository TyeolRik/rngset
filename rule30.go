package rngset

import "fmt"

// I couldn't find exact document which is about initial seeds.
// Some paper said, initial seed is "Skipping n-bits and pick batches of n" (https://dev.to/arpit_bhayani/pseudorandom-numbers-using-rule-30-4mif)
// But, there are some uncertainties. (What is "Standard initial dummy bits", which means before skipping.)
// So, I utilized a demonstrated website (Wolfram Cloud Demonstration Page)
// https://demonstrations.wolfram.com/UsingRule30ToGeneratePseudorandomRealNumbers/

// There could be error if seed is zero.
// So, if seed is zero, seed changed to my birthday (617) on my own authority
func (r *rule30) setSeed(seed []uint8) {
	var checker uint8 = 0
	startPoint := (r.width - int64(len(seed))) / 2
	for i, bit := range seed {
		r.board[0][startPoint+int64(i)] = bit
		checker += bit
	}
	r.seed = seed
	if checker == 0 {
		new := NewRule30(617, r.height)
		r = &new
	}
}

func (r *rule30) refresh() {
	for col := 1; col < int(r.height); col++ {
		for row := 0; row < int(r.width); row++ {
			if row == 0 {
				r.board[col][row] = 0 ^ (r.board[col-1][row] | r.board[col-1][row+1])
			} else if row == int(r.width)-1 {
				r.board[col][row] = r.board[col-1][row-1] ^ r.board[col-1][row]
			} else {
				r.board[col][row] = r.board[col-1][row-1] ^ (r.board[col-1][row] | r.board[col-1][row+1])
			}
		}
	}
}

func (r *rule30) Generate() (ret float64) {
	div := 0.5
	newSeedArray := make([]uint8, r.height)
	for col := 0; col < int(r.height); col++ {
		ret += float64(r.board[col][r.height]) * div
		newSeedArray[col] = r.board[col][r.height-1+int64(len(r.seed)/2)]
		div = div / 2.0
	}
	r.setSeed(newSeedArray)
	r.refresh()
	return
}

func (r *rule30) Show() {
	for col := range r.board {
		for row := range r.board[col] {
			fmt.Printf("%d ", r.board[col][row])
		}
		fmt.Printf("\n")
	}
}
