package base

// Board represents a general current state.
// NOTE: should be comparable in map key types, as it is used in ../algorithms.
type Board interface {
	String() string
	// Equals checks the equality to the other board.
	Equals(other Board) bool
	// Operate returns a new board after applying the given operation.
	// Expects NOT to modify the callee.
	Operate(op Operation) Board
	// PossibleOps returns possible operations from this state.
	PossibleOps() []Operation
	// IsGoal checks if this state is the goal state.
	IsGoal() bool
	// IsSolvable checks if the goal state is at all reachable from this state,
	// if it can be calculated.
	IsSolvable() bool
}

// Operation represents a general operation on a state.
type Operation interface {
	String() string
}
