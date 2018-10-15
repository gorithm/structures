package list

// Iterator iterates over the elements in a list.
type Iterator interface {
	HasNext() bool
	HasPrevious() bool
	Next() interface{}
	NextIndex() int
	Previous() interface{}
	PreviousIndex() int
}
