package linkedlist

// LinkedList implements a doubly-linked list.
//
// LinkedList implements the List and Collection interfaces.
type (
	LinkedList struct {
		head *node
		tail *node
		size int
	}

	node struct {
		value interface{}
		prev  *node
		next  *node
	}
)

func newNode(v interface{}) *node {
	return &node{
		value: v,
		prev:  nil,
		next:  nil,
	}
}

// Add adds an element to the end of the linked list.
func (l *LinkedList) Add(v interface{}) {
	n := newNode(v)
	if l.tail != nil {
		n.prev = l.tail
		l.tail.next = n
	} else {
		l.head = n
	}
	l.tail = n
	l.size++
}

// AddFirst adds an element to the beginning of the linked list.
func (l *LinkedList) AddFirst(v interface{}) {
	n := newNode(v)
	if l.head != nil {
		n.next = l.head
		l.head.prev = n
	} else {
		l.tail = n
	}
	l.head = n
	l.size++
}

// GetFirst returns the element at the beginning of the list.
func (l *LinkedList) GetFirst() (interface{}, error) {
	if l.IsEmpty() {
		return nil, ErrEmptyLinkedList
	}
	return l.head.value, nil
}

// GetLast returns the element at the end of the linked list.
func (l *LinkedList) GetLast() (interface{}, error) {
	if l.IsEmpty() {
		return nil, ErrEmptyLinkedList
	}
	return l.tail.value, nil
}

// RemoveFirst removes the element at the beginning of the linked list.
func (l *LinkedList) RemoveFirst() error {
	if l.IsEmpty() {
		return ErrEmptyLinkedList
	}
	l.head = l.head.next
	l.size--
	if l.Size() == 0 {
		l.tail = nil
	}
	if l.head != nil {
		l.head.prev = nil
	}
	return nil
}

// RemoveLast removes the element at the end of the linked list.
func (l *LinkedList) RemoveLast() error {
	if l.IsEmpty() {
		return ErrEmptyLinkedList
	}
	l.tail = l.tail.prev
	l.size--
	if l.Size() == 0 {
		l.head = nil
	}
	if l.tail != nil {
		l.tail.next = nil
	}
	return nil
}

// Size returns the number of elements in the linked list.
func (l *LinkedList) Size() int {
	return l.size
}

// IsEmpty returns whether or not the list is empty.
func (l *LinkedList) IsEmpty() bool {
	return l.Size() == 0
}

// Get returns the element at a given index in the linked list.
func (l *LinkedList) Get(i int) (interface{}, error) {
	// TODO: Implement Get.
	return nil, nil
}

// TODO: Implement LinkedList

// Collection methods:
//
// Contains(v interface{}) bool
// ContainsAll(c Collection) bool
// Equals(v interface{}) bool
// IsEmpty() bool <- Implemented.
// Iterator() Iterator
// Size() int <- Implemented.
// Array() []interface{}

// List methods:
//
// Get(i int) (interface{}, error)
// IndexOf(v interface{}) (int, error)
// LastIndexOf(v interface{}) (int, error)
// ReplaceAll(func(v interface{}) interface{})
// Sort() error
// SortWithComparator(ctr comparator.Comparator) error
// SubList(beg, end int) (List, error)
