package list

import "github.com/gorithm/structures/collection/comparator"

// List organizes data in a sequential, ordered manner.
type List interface {
	Get(i int) (interface{}, error)
	IndexOf(v interface{}) (int, error)
	LastIndexOf(v interface{}) (int, error)
	ReplaceAll(func(v interface{}) interface{})
	Sort() error
	SortWithComparator(ctr comparator.Comparator) error
	SubList(beg, end int) (List, error)
}
