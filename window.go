package window

// Contains reports whether value belongs to the inclusive interval [first, last].
// An interval whose first bound exceeds its last bound contains no values.
func Contains(first, last, value int) bool {
	return first <= value && value <= last
}
