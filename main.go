package main

import (
	"./_16puzzle"
	"./algorithms/a_star"
	"./algorithms/best_first"
	"./algorithms/bfs"
	"./algorithms/ida"
	"./algorithms/iddfs"
	"./base"
	"fmt"
	"time"
)

const repeat = 3

func main() {
	board := _16puzzle.GetExample()

	type solveFunc func(b base.Board) ([]base.Operation, time.Duration)
	type kind struct {
		f    solveFunc
		name string
	}

	kinds := [11]kind{
		{bfs.Solve, "Breadth-First Search"},
		{iddfs.Solve, "Iterative-Deepening Depth-First Search"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return a_star.Solve(b, _16puzzle.H1)
		}, "A* Search (H1)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return a_star.Solve(b, _16puzzle.H2)
		}, "A* Search (H2)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return a_star.Solve(b, _16puzzle.H3)
		}, "A* Search (H3)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return ida.Solve(b, _16puzzle.H1)
		}, "IDA* Search (H1)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return ida.Solve(b, _16puzzle.H2)
		}, "IDA* Search (H2)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return ida.Solve(b, _16puzzle.H3)
		}, "IDA* Search (H3)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return best_first.Solve(b, _16puzzle.H1)
		}, "Best First Search (H1)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return best_first.Solve(b, _16puzzle.H2)
		}, "Best First Search (H2)"},
		{func(b base.Board) ([]base.Operation, time.Duration) {
			return best_first.Solve(b, _16puzzle.H3)
		}, "Best First Search (H3)"},
	}
	solve := [11]bool{
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
		true,
	}
	times := [11]time.Duration{}

	for i := 0; i < 11; i++ {
		if !solve[i] {
			continue
		}
		fmt.Println("---- " + kinds[i].name + " ----")
		for r := 0; r < repeat; r++ {
			fmt.Printf("-- %v-th repeat\n", r+1)
			ops, duration := kinds[i].f(board)
			times[i] += duration

			for i, layer := range ops {
				fmt.Printf("%v-th move: %v\n", i+1, layer)
			}

			fmt.Printf("-- time: %v\n", duration)
		}
		fmt.Printf("--- average time: %v\n", times[i]/repeat)
		fmt.Println()
		fmt.Println()
	}
}
