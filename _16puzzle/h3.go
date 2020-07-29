package _16puzzle

import (
	"../base"
	"fmt"
)

const goalTableLimit = 15

var (
	goalTable map[board16Puzzle]int
)

func init() {
	goalTable = genGoalTable(goalTableLimit)
	fmt.Printf("generated exact steps table from the goal, size %d\n", len(goalTable))
}

type boardSearch struct {
	b     board16Puzzle
	depth int
}

// generates exact steps table from the goal using BFS
func genGoalTable(limit int) map[board16Puzzle]int {
	ret := make(map[board16Puzzle]int)

	next := make([]*boardSearch, 0)
	next = append(next, &boardSearch{
		b:     Goal,
		depth: 0,
	})

	for len(next) > 0 {
		bs := next[0]
		next = next[1:]
		// check visited
		if _, ok := ret[bs.b]; ok {
			continue
		}
		ret[bs.b] = bs.depth

		if bs.depth >= limit {
			continue
		}

		for _, op := range bs.b.PossibleOps() {
			nextBoard := bs.b.Operate(op).(board16Puzzle)
			next = append(next, &boardSearch{
				b:     nextBoard,
				depth: bs.depth + 1,
			})
		}
	}

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func H3(b base.Board) int {
	board, ok := b.(board16Puzzle)
	if !ok {
		panic("not a 16 puzzle board")
	}

	h2 := table[genRowWDTable(&board)] + table[genColWDTable(&board)]
	supplement, ok := goalTable[board]
	if !ok {
		supplement = goalTableLimit + 1
	}
	return max(h2, supplement)
}
