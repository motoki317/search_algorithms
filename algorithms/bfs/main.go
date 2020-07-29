package bfs

import (
	"../../base"
	"fmt"
	"time"
)

type Searcher struct {
	next         []*BoardSearch
	visited      map[base.Board]bool
	currentDepth int
	// searched nodes count
	nodes int
}

// for displaying operations with BFS
// board type that has the
type BoardSearch struct {
	current   base.Board
	before    *BoardSearch
	operation base.Operation
	// cache
	heuristics int
	depth      int
}

func (b *BoardSearch) restore() []base.Operation {
	ret := make([]base.Operation, 0)
	cur := b
	for cur != nil {
		ret = append([]base.Operation{cur.operation}, ret...)
		cur = cur.before
	}
	return ret[1:]
}

// search executes BFS, returning non-nil stack if the goal was found.
func (a *Searcher) search(first base.Board) []base.Operation {
	a.next = append(a.next, &BoardSearch{
		current: first,
		before:  nil,
		// to be cut in restore()
		operation: nil,
		depth:     0,
	})

	for len(a.next) > 0 {
		bs := a.next[0]
		a.next = a.next[1:]
		// check visited
		if a.visited[bs.current] {
			continue
		}
		a.visited[bs.current] = true
		a.nodes++

		// print current depth
		if a.currentDepth < bs.depth {
			fmt.Printf("searching depth %v\n", bs.depth)
			a.currentDepth = bs.depth
		}

		if bs.current.IsGoal() {
			return bs.restore()
		}

		for _, op := range bs.current.PossibleOps() {
			next := bs.current.Operate(op)
			a.next = append(a.next, &BoardSearch{
				current:   next,
				before:    bs,
				operation: op,
				depth:     bs.depth + 1,
			})
		}
	}

	return nil
}

func NewSearcher() *Searcher {
	return &Searcher{
		next:    make([]*BoardSearch, 0),
		visited: make(map[base.Board]bool),
	}
}

func Solve(b base.Board) (operations []base.Operation, duration time.Duration) {
	if !b.IsSolvable() {
		fmt.Println("Not solvable")
		return
	}

	s := NewSearcher()

	start := time.Now()
	operations = s.search(b)
	if operations == nil {
		fmt.Println("No goal found")
		return
	}
	duration = time.Since(start)
	fmt.Printf("Searched nodes: %v\n", s.nodes)
	fmt.Printf("Remaining nodes: %v\n", len(s.next))
	return
}
