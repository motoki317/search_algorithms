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

func iddfs(b base.Board, depth, maxDepth int) ([]base.Operation, bool) {
	nodes++

	if b.IsGoal() {
		return nil, true
	}

	if depth >= maxDepth {
		return nil, false
	}

	if beforeDepth, ok := visited[b]; ok && beforeDepth <= depth {
		return nil, false
	}
	visited[b] = depth

	for _, op := range b.PossibleOps() {
		next := b.Operate(op)
		ops, res := iddfs(next, depth+1, maxDepth)
		if res {
			ops = append(ops, op)
			return ops, true
		}
	}
	return nil, false
}

func reverse(s []base.Operation) []base.Operation {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func Solve(b base.Board) (operations []base.Operation, duration time.Duration) {
	if !b.IsSolvable() {
		fmt.Println("Not solvable")
		return
	}

	nodes = 0
	start := time.Now()
	for i := 1; i <= 100; i++ {
		fmt.Printf("max depth: %v\n", i)
		visited = make(map[base.Board]int)
		var ok bool
		if operations, ok = iddfs(b, 0, i); ok {
			operations = reverse(operations)
			duration = time.Since(start)
			fmt.Printf("Goal found at depth %v\n", i)
			fmt.Printf("Searched nodes: %v\n", nodes)
			return
		}
	}
	fmt.Println("No goal found")
	return
}
