package stack

import "errors"

// Errors that can be produced by the Stack struct.
var (
	ErrEmptyStack = errors.New("gorithm/stack: stack is empty")
)
