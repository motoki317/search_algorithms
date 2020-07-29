package _16puzzle

import (
	"../base"
	"fmt"
)

var table map[wdTable]int

func init() {
	table = generateWDDatabase()
}

// ref: http://www.ic-net.or.jp/home/takaken/nt/slide/solve15.html
// For Heuristics function 'Walking Distance' or 'WD'

type wdTable [size][size]int

func genRowWDTable(b *board16Puzzle) wdTable {
	var ret wdTable
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if b[row][col] == 0 {
				continue
			}
			// which row this number should belong to
			toRow := (b[row][col] - 1) / size
			ret[row][toRow]++
		}
	}
	return ret
}

func genColWDTable(b *board16Puzzle) wdTable {
	var ret wdTable
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if b[row][col] == 0 {
				continue
			}
			// which column this number should belong to
			toCol := (b[row][col] - 1) % size
			ret[col][toCol]++
		}
	}
	return ret
}

func startingWDTable() wdTable {
	var ret wdTable
	for i := 0; i < size-1; i++ {
		ret[i][i] = size
	}
	ret[size-1][size-1] = size - 1
	return ret
}

// returns 0-indexed row number in which there is a space.
func (w wdTable) zeroRow() int {
	for i := 0; i < size; i++ {
		sum := 0
		for j := 0; j < size; j++ {
			sum += w[i][j]
		}
		if sum == size-1 {
			return i
		}
	}
	panic(fmt.Sprintf("zero row not found with table %v\n", w))
}

func (w wdTable) next() []wdTable {
	zeroRow := w.zeroRow()
	ret := make([]wdTable, 0)
	if zeroRow > 0 {
		for j := 0; j < size; j++ {
			if w[zeroRow-1][j] > 0 {
				next := w
				next[zeroRow-1][j]--
				next[zeroRow][j]++
				ret = append(ret, next)
			}
		}
	}
	if zeroRow < size-1 {
		for j := 0; j < size; j++ {
			if w[zeroRow+1][j] > 0 {
				next := w
				next[zeroRow+1][j]--
				next[zeroRow][j]++
				ret = append(ret, next)
			}
		}
	}
	return ret
}

func generateWDDatabase() map[wdTable]int {
	// to save depth
	type WDTableDepth struct {
		w     wdTable
		depth int
	}
	ret := make(map[wdTable]int)
	queue := make([]WDTableDepth, 0)
	queue = append(queue, WDTableDepth{
		w:     startingWDTable(),
		depth: 0,
	})
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if _, ok := ret[current.w]; ok {
			continue
		}
		ret[current.w] = current.depth

		for _, next := range current.w.next() {
			if _, ok := ret[next]; ok {
				continue
			}
			queue = append(queue, WDTableDepth{
				w:     next,
				depth: current.depth + 1,
			})
		}
	}

	fmt.Printf("Generated WD table with size %v\n", len(ret))
	return ret
}

func H2(b base.Board) int {
	board, ok := b.(board16Puzzle)
	if !ok {
		panic("not a 16 puzzle board")
	}
	return table[genRowWDTable(&board)] + table[genColWDTable(&board)]
}
