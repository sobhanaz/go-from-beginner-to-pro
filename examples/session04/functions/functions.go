// Session 04 — defining and calling functions.
// Run:  go run examples/session04/functions/functions.go
package main

import "fmt"

// A function with two int parameters that returns an int.
func add(a int, b int) int {
	return a + b
}

// Shorthand: consecutive same-type params share one type annotation.
func multiply(a, b int) int {
	return a * b
}

// A function that returns nothing (no return type after the parens).
func greet(name string) {
	fmt.Println("Hello,", name)
}

// Functions can call other functions. Order of definition does not matter.
func areaOfSquare(side int) int {
	return multiply(side, side)
}

func main() {
	sum := add(3, 4)
	fmt.Println("3 + 4 =", sum)

	fmt.Println("6 * 7 =", multiply(6, 7))

	greet("Sobhan")

	fmt.Println("area of 5x5 square:", areaOfSquare(5))
}
