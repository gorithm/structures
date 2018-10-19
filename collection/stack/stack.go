package stack

// Stack is a collection which implements a LIFO data access schema.
type Stack interface {
	Push(v interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
}
