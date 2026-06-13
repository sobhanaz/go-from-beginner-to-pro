// Session 06 — fixed-size arrays.
// Run:  go run examples/session06/arrays/arrays.go
package main

import "fmt"

func main() {
	// Declared with a fixed size; starts zero-valued.
	var nums [3]int
	nums[0] = 10
	nums[1] = 20
	fmt.Println("nums:", nums, "len:", len(nums)) // [10 20 0] len: 3

	// Declare with values.
	primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println("primes:", primes)

	// Let Go count the length with [...].
	days := [...]string{"Mon", "Tue", "Wed"}
	fmt.Println("days:", days, "len:", len(days))

	// Arrays are VALUE types: assigning copies the whole array.
	a := [3]int{1, 2, 3}
	b := a    // full copy
	b[0] = 99 // changing b does NOT affect a
	fmt.Println("a:", a, "b:", b) // a:[1 2 3] b:[99 2 3]
}
