package a_star

import (
	"../../base"
	"../heuristics"
	"fmt"
	"github.com/emirpasic/gods/trees/binaryheap"
	"time"
)

type AStarSearcher struct {
	next    *binaryheap.Heap
	visited map[base.Board]bool
	// visited nodes count
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

// search executes A* search, returning non-nil stack if the goal was found.
func (a *AStarSearcher) search(first base.Board) []base.Operation {
	a.next.Push(&BoardSearch{
		current: first,
		before:  nil,
		// to be cut in restore()
		operation: nil,
		depth:     0,
	})

	for !a.next.Empty() {
		c, _ := a.next.Pop()
		bs := c.(*BoardSearch)
		// check visited
		if a.visited[bs.current] {
			continue
		}
		a.visited[bs.current] = true
		a.nodes++

		//fmt.Printf("Searching node with h(X) = %v\n", bs.current.heuristics())
		if bs.current.IsGoal() {
			return bs.restore()
		}

		for _, op := range bs.current.PossibleOps() {
			next := (bs.current).Operate(op)
			a.next.Push(&BoardSearch{
				current:   next,
				before:    bs,
				operation: op,
				depth:     bs.depth + 1,
			})
		}
	}

	return nil
}

func NewSearcher(h heuristics.F) *AStarSearcher {
	return &AStarSearcher{
		next: binaryheap.NewWith(func(a, b interface{}) int {
			aa, bb := a.(*BoardSearch), b.(*BoardSearch)
			if aa.heuristics == 0 {
				aa.heuristics = h(aa.current)
			}
			if bb.heuristics == 0 {
				bb.heuristics = h(bb.current)
			}
			return (aa.depth + aa.heuristics) - (bb.depth + bb.heuristics)
		}),
		visited: make(map[base.Board]bool),
	}
}

func Solve(b base.Board, h heuristics.F) (duration time.Duration) {
	if !b.IsSolvable() {
		fmt.Println("Not solvable")
		return
	}
	fmt.Printf("heuristics: %v\n", h(b))

	s := NewSearcher(h)

	start := time.Now()
	res := s.search(b)
	if res == nil {
		fmt.Println("No goal found")
		return
	}
	duration = time.Since(start)
	fmt.Printf("Searched nodes: %v\n", s.nodes)
	fmt.Printf("Remaining nodes: %v\n", s.next.Size())

	for i, layer := range res {
		fmt.Printf("%v-th move: %v\n", i+1, layer)
	}
	return
}
