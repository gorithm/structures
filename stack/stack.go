package stack

type (
	// Stack is a type-agnostic implementation of a LIFO data structure.
	Stack struct {
		top  *node
		size int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// New returns a new pointer to an empty Stack.
func New() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

// Push inserts a new value into the stack.
func (s *Stack) Push(val interface{}) {
	s.top = &node{
		value: val,
		prev:  s.top,
	}
	s.size++
}

// Peek returns the value of the top of the Stack.
func (s *Stack) Peek() (interface{}, error) {
	if s.Size() == 0 {
		return nil, ErrEmptyStack
	}
	val := s.top.value
	return val, nil
}

// Pop returns the value of the top of the Stack and removes the node.
func (s *Stack) Pop() (interface{}, error) {
	val, err := s.Peek()
	if err != nil {
		return nil, err
	}
	s.top = s.top.prev
	s.size--
	return val, nil
}

// Size returns the number of elements in the Stack.
func (s *Stack) Size() int {
	return s.size
}
