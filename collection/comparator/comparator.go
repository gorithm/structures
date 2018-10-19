package comparator

// Comparator is a function that compares two interfaces.
//
// The comparator should be implemented in the following form:
// If a > b, return > 0
// If a < b, return < 0
// If a == b, return 0.
type Comparator func(a, b interface{}) int
