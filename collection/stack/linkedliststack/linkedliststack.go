package linkedliststack

import "github.com/gorithm/structures/collection/list/linkedlist"

// LinkedListStack is a stack that is implemented with a collection.LinkedList underneath.
//
// LinkedListStack implements the Stack and Collection interfaces.
type LinkedListStack struct {
	l *linkedlist.LinkedList
}

// New returns a pointer to a new linked list stack
func New() *LinkedListStack {
	return &LinkedListStack{
		l: linkedlist.New(),
	}
}

// Push adds an element to the top of the stack.
func (s *LinkedListStack) Push(v interface{}) {
	s.l.AddFirst(v)
}

// Peek returns the top element of the stack.
func (s *LinkedListStack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}
	v, err := s.l.GetFirst()
	if err != nil {
		return nil, err
	}
	return v, nil
}

// Pop returns and removes the top element of the stack.
func (s *LinkedListStack) Pop() (interface{}, error) {
	v, err := s.Peek()
	if err != nil {
		return nil, err
	}
	if err := s.l.RemoveFirst(); err != nil {
		return nil, err
	}
	return v, nil
}

// Size returns the number of elements in the stack.
func (s *LinkedListStack) Size() int {
	return s.l.Size()
}

// IsEmpty returns whether or not the stack is empty.
func (s *LinkedListStack) IsEmpty() bool {
	return s.Size() == 0
}
