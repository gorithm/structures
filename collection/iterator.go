package collection

// Iterator iterates through the values in any kind of collection.
type Iterator interface {
	HasNext() bool
	Next() (interface{}, error)
}
