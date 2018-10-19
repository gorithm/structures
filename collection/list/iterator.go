package list

// Iterator iterates over the elements in a list.
type Iterator interface {
	HasNext() bool
	HasPrevious() bool
	Next() (interface{}, error)
	NextIndex() (int, error)
	Previous() (interface{}, error)
	PreviousIndex() (int, error)
}
