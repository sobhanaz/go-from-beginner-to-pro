// Session 05 — functions as values, and closures.
// Run:  go run examples/session05/closures/closures.go
package main

import "fmt"

// counter returns a function that remembers its own count.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// multiplier returns a function that multiplies by a captured factor.
func multiplier(factor int) func(int) int {
	return func(n int) int {
		return n * factor
	}
}

func main() {
	// A function stored in a variable (anonymous function).
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println("add(2,3) =", add(2, 3))

	// Two independent counters — each closes over its own `count`.
	c1 := counter()
	c2 := counter()
	fmt.Println("c1:", c1(), c1(), c1()) // 1 2 3
	fmt.Println("c2:", c2())             // 1 (separate state)

	// Build specialized functions from a closure.
	double := multiplier(2)
	triple := multiplier(3)
	fmt.Println("double(10) =", double(10)) // 20
	fmt.Println("triple(10) =", triple(10)) // 30
}
