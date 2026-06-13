// Session 16 — the code under test. A tiny calculator package.
// Note: this is a LIBRARY package (not "main"), so it has no func main().
package calc

import "errors"

// ErrDivideByZero is returned by Divide when the divisor is zero.
var ErrDivideByZero = errors.New("cannot divide by zero")

// Add returns a + b.
func Add(a, b int) int {
	return a + b
}

// Divide returns a / b, or an error if b is zero.
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrDivideByZero
	}
	return a / b, nil
}

// IsEven reports whether n is even.
func IsEven(n int) bool {
	return n%2 == 0
}
