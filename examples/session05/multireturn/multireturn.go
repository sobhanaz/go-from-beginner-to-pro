// Session 05 — multiple return values and the value,err pattern.
// Run:  go run examples/session05/multireturn/multireturn.go
package main

import (
	"errors"
	"fmt"
	"math"
)

// Returns two values: quotient and remainder.
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// The idiomatic value,err pattern: result + an error.
// A nil error means success.
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("cannot sqrt a negative number")
	}
	return math.Sqrt(x), nil
}

// Named return values + a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // sends back x and y
}

func main() {
	q, r := divide(17, 5)
	fmt.Printf("17 / 5 = %d remainder %d\n", q, r)

	// Discard a return you don't need with _.
	_, remainder := divide(17, 5)
	fmt.Println("just the remainder:", remainder)

	// Handle the error: the pattern you'll write thousands of times.
	if result, err := sqrt(16); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("sqrt(16) =", result)
	}

	if _, err := sqrt(-4); err != nil {
		fmt.Println("sqrt(-4) failed:", err)
	}

	a, b := split(17)
	fmt.Printf("split(17) -> %d, %d\n", a, b)
}
