package collection

// Collection provides a uniform way of accessing data from lists, plots, and sets.
type Collection interface {
	Contains(v interface{}) bool
	ContainsAll(c Collection) bool
	Equals(v interface{}) bool
	IsEmpty() bool
	Iterator() Iterator
	Size() int
	Array() []interface{}
}
