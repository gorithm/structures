package arraylist

// ArrayList implements Go's slices in the back-end, but allows for extra functionality
// provided included with the struct's methods.
//
// ArrayList implements the List and Collection interfaces.
type ArrayList struct {
	arr []interface{}
}

// IsEmpty returns whether or not the ArrayList is empty.
func (a *ArrayList) IsEmpty() bool {
	return len(a.arr) == 0
}

// Add adds an element to the end of the ArrayList.
func (a *ArrayList) Add(v interface{}) {
	a.arr = append(a.arr, v)
}

// TODO: Implement ArrayList

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
