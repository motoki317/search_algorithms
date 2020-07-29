package _16puzzle

import (
	"../base"
	"fmt"
	"strconv"
	"strings"
)

type (
	board16Puzzle     [size][size]int
	operation16Puzzle int
)

var Goal board16Puzzle

const size = 4

const (
	UP operation16Puzzle = iota
	DOWN
	LEFT
	RIGHT
)

func init() {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			Goal[i][j] = i*size + j + 1
		}
	}
	Goal[size-1][size-1] = 0
}

func newBoard(data [size][size]int) board16Puzzle {
	return data
}

var operations = []operation16Puzzle{UP, DOWN, LEFT, RIGHT}

func (b board16Puzzle) String() string {
	ret := "----\n"
	for i := 0; i < size; i++ {
		str := make([]string, 0, size)
		for j := 0; j < size; j++ {
			str = append(str, strconv.Itoa(b[i][j]))
		}
		ret += fmt.Sprintf("| %v |\n", strings.Join(str, " | "))
	}
	ret += "----"
	return ret
}

func (b board16Puzzle) Equals(other base.Board) bool {
	otherBoard, ok := other.(board16Puzzle)
	if !ok {
		return false
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b[i][j] != otherBoard[i][j] {
				return false
			}
		}
	}
	return true
}

func (b board16Puzzle) zeroPos() (int, int) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if b[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (b board16Puzzle) opIsPossible(op base.Operation) bool {
	i, j := b.zeroPos()
	if i == -1 && j == -1 {
		return false
	}
	switch op {
	case UP:
		return i != size-1
	case DOWN:
		return i != 0
	case LEFT:
		return j != size-1
	case RIGHT:
		return j != 0
	}
	return false
}

func (op operation16Puzzle) String() string {
	switch op {
	case UP:
		return "UP ↑"
	case DOWN:
		return "DOWN ↓"
	case LEFT:
		return "LEFT ←"
	case RIGHT:
		return "RIGHT →"
	}
	return "unknown"
}

func (b board16Puzzle) Operate(op base.Operation) base.Board {
	i, j := b.zeroPos()
	switch op {
	case UP:
		b[i][j] = b[i+1][j]
		b[i+1][j] = 0
	case DOWN:
		b[i][j] = b[i-1][j]
		b[i-1][j] = 0
	case LEFT:
		b[i][j] = b[i][j+1]
		b[i][j+1] = 0
	case RIGHT:
		b[i][j] = b[i][j-1]
		b[i][j-1] = 0
	}
	return b
}

func (b board16Puzzle) PossibleOps() []base.Operation {
	ret := make([]base.Operation, 0)
	for _, op := range operations {
		if !b.opIsPossible(op) {
			continue
		}
		ret = append(ret, op)
	}
	return ret
}

func (b board16Puzzle) IsGoal() bool {
	return b.Equals(Goal)
}

func (b board16Puzzle) parity() bool {
	ret := 0
	visited := make([]bool, size*size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			next := b[i][j]
			cycleLen := 0
			for !visited[next] {
				visited[next] = true
				cycleLen++
				next = b[next/size][next%size]
			}
			if cycleLen == 0 {
				continue
			}
			ret += cycleLen - 1
		}
	}
	return ret%2 == 0
}

func (b board16Puzzle) puzzleParity() bool {
	p := b.parity()
	i, j := b.zeroPos()
	return p != ((i+j)%2 == 0)
}

func (b board16Puzzle) IsSolvable() bool {
	return b.puzzleParity() == Goal.puzzleParity()
}
