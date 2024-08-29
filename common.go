package essential

// Ternary operator
//
// Returns a if c is true, b otherwise
func Ternary[T any](c bool, a T, b T) T {
	if c {
		return a
	}

	return b
}
