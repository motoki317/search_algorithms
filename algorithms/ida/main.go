package ida

import (
	"../../base"
	"../heuristics"
	"fmt"
	"github.com/emirpasic/gods/stacks"
	"github.com/emirpasic/gods/stacks/arraystack"
	"time"
)

type AStarSearcher struct {
	// visited nodes count
	totalNodes   int
	currentNodes int
	// heuristics function
	h heuristics.F
}

const INF = 1e9

// ref: https://en.wikipedia.org/wiki/Iterative_deepening_A*

func pathContains(path stacks.Stack, target base.Board) bool {
	for _, v := range path.Values() {
		b := v.(base.Board)
		if b.Equals(target) {
			return true
		}
	}
	return false
}

func (a *AStarSearcher) idaSearch(path stacks.Stack, depth int, limit int) (found bool, nextLimit int) {
	l, _ := path.Peek()
	current := l.(base.Board)

	a.currentNodes++

	f := depth + a.h(current)
	if f > limit {
		return false, f
	}
	if current.IsGoal() {
		return true, 0
	}

	// minimum cost (f) above the current limit
	nextLimit = INF
	for _, op := range current.PossibleOps() {
		nextBoard := current.Operate(op)
		if pathContains(path, nextBoard) {
			continue
		}
		path.Push(nextBoard)
		found, lim := a.idaSearch(path, depth+1, limit)
		if found {
			return found, 0
		}
		if lim < nextLimit {
			nextLimit = lim
		}
		path.Pop()
	}
	return false, nextLimit
}

// executes iterative deepening A* search (IDA*).
func (a *AStarSearcher) ida(root base.Board) (path stacks.Stack, limit int) {
	limit = a.h(root)
	path = arraystack.New()
	path.Push(root)

	defer func() {
		fmt.Printf("Searched %v nodes\n", a.currentNodes)
		a.totalNodes += a.currentNodes
	}()
	for {
		fmt.Printf("Searching with limit %v\n", limit)
		found, nextLimit := a.idaSearch(path, 0, limit)
		if found {
			return path, limit
		}
		if nextLimit == INF {
			break
		}
		limit = nextLimit
		fmt.Printf("Searched %v nodes\n", a.currentNodes)
		a.totalNodes += a.currentNodes
		a.currentNodes = 0
	}
	return nil, 0
}

func NewSearcher(h heuristics.F) *AStarSearcher {
	return &AStarSearcher{
		h: h,
	}
}

func reverse(s []interface{}) []interface{} {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func restoreOperation(before, after base.Board) base.Operation {
	for _, op := range before.PossibleOps() {
		operated := before.Operate(op)
		if operated.Equals(after) {
			return op
		}
	}
	panic("cannot find operation")
}

func Solve(b base.Board, h heuristics.F) (operations []base.Operation, duration time.Duration) {
	if !b.IsSolvable() {
		fmt.Println("Not solvable")
		return
	}

	s := NewSearcher(h)
	fmt.Printf("heuristics: %v\n", s.h(b))

	start := time.Now()
	path, limit := s.ida(b)
	if path == nil {
		fmt.Println("No goal found")
		return
	}
	duration = time.Since(start)
	fmt.Printf("Searched total nodes: %v\n", s.totalNodes)
	fmt.Printf("Searched nodes last time (for calculating effective branching factor): %v\n", s.currentNodes)
	fmt.Printf("Goal found at limit %v\n", limit)

	//fmt.Printf("first state:\n%v\n", b)
	values := reverse(path.Values())
	operations = make([]base.Operation, 0, len(values)-1)
	for i := range values {
		if i == 0 {
			continue
		}
		before := values[i-1].(base.Board)
		after := values[i].(base.Board)
		op := restoreOperation(before, after)
		operations = append(operations, op)
	}
	return
}
