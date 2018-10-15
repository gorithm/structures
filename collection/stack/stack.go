package stack

// Stack is a collection which implements a LIFO data access schema.
type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}
