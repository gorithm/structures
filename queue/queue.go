package queue

type (
	// Queue is a type-agnostic implementation of a FIFO data structure.
	Queue struct {
		first *node
		last  *node
		size  int
	}

	node struct {
		value interface{}
		next  *node
	}
)

// New returns a new pointer to an empty Queue.
func New() *Queue {
	return &Queue{
		first: nil,
		last:  nil,
		size:  0,
	}
}

// Size returns the number of elements in the Queue.
func (q *Queue) Size() int {
	return q.size
}

// Enqueue adds an new element to the end of the Queue.
func (q *Queue) Enqueue(v interface{}) {
	n := node{
		value: v,
		next:  nil,
	}
	if q.Size() == 0 {
		q.first = &n
		q.last = q.first
	} else {
		q.last.next = &n
		q.last = q.last.next
	}
	q.size++
}

// Peek returns the first element in the queue
func (q *Queue) Peek() (interface{}, error) {
	if q.Size() == 0 {
		return nil, ErrEmptyQueue
	}
	return q.first.value, nil
}

// Dequeue returns and removes the first element in the Queue.
func (q *Queue) Dequeue() (interface{}, error) {
	val, err := q.Peek()
	if err != nil {
		return nil, err
	}
	if q.Size() == 1 {
		q.last = nil
	}
	q.first = q.first.next
	q.size--
	return val, nil
}
