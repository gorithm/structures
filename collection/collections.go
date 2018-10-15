package collection

// Collection provides a class of operations to manipulate lists, plots (maps), and sets.
type Collection interface {
	Contains(v interface{}) bool
	ContainsAll(c Collection) bool
	Equals(v interface{}) bool
	IsEmpty() bool
	Iterator() Iterator
	Size() int
	Array() []interface{}
}
