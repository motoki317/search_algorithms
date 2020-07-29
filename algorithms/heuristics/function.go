package heuristics

import "../../base"

// heuristics.F represents a generic heuristics function,
// which returns speculative distance to the goal state.
type F func(b base.Board) int
