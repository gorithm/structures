package queue

import "errors"

// Errors that can be produced by the Queue struct.
var (
	ErrEmptyQueue = errors.New("gorithm/queue: queue is empty")
)
