package queue

// Queue is a collection designed for holding elements prior to processing.
// Besides basic Collection operations, queues provide additional insertion, extraction, and inspection operations.
// Each of these methods exists in two forms: one throws an exception if the operation fails,
// the other returns a special value (either null or false, depending on the operation).
// The latter form of the insert operation is designed specifically for use with capacity-restricted Queue implementations;
// in most implementations, insert operations cannot fail.
type Queue interface {
	Enqueue(v interface{})
	Dequeue() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}
