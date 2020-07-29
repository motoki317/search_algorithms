package _16puzzle

import (
	"../base"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Sum of Manhattan distance
func H1(b base.Board) int {
	board, ok := b.(board16Puzzle)
	if !ok {
		panic("not a 16 puzzle board")
	}
	ret := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] == 0 {
				continue
			}
			num := board[i][j]
			// goal index
			ii := (num - 1) / size
			jj := (num - 1) % size
			ret += abs(i-ii) + abs(j-jj)
		}
	}
	return ret
}
