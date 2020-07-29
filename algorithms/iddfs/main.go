package iddfs

import (
	"../../base"
	"fmt"
	"time"
)

var (
	nodes   = 0
	visited = make(map[base.Board]int)
)

func iddfs(b base.Board, depth, maxDepth int) bool {
	nodes++

	if b.IsGoal() {
		return true
	}

	if depth >= maxDepth {
		return false
	}

	if beforeDepth, ok := visited[b]; ok && beforeDepth <= depth {
		return false
	}
	visited[b] = depth

	for _, op := range b.PossibleOps() {
		next := b.Operate(op)
		res := iddfs(next, depth+1, maxDepth)
		if res {
			fmt.Printf("depth %v operation: %v\n", depth, op)
			return true
		}
	}
	return false
}

func Solve(b base.Board) (duration time.Duration) {
	if !b.IsSolvable() {
		fmt.Println("Not solvable")
		return
	}

	nodes = 0
	start := time.Now()
	for i := 1; i <= 100; i++ {
		fmt.Printf("max depth: %v\n", i)
		visited = make(map[base.Board]int)
		if iddfs(b, 0, i) {
			duration = time.Since(start)
			fmt.Printf("Goal found at depth %v\n", i)
			fmt.Printf("Searched nodes: %v\n", nodes)
			return
		}
	}
	fmt.Println("No goal found")
	return
}
