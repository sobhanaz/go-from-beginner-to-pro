// Session 05 — variadic functions (...) accept any number of args.
// Run:  go run examples/session05/variadic/variadic.go
package main

import "fmt"

// `numbers` is a slice of int inside the function.
func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	fmt.Println("sum(1,2,3)       =", sum(1, 2, 3))
	fmt.Println("sum(10,20,30,40) =", sum(10, 20, 30, 40))
	fmt.Println("sum()            =", sum()) // zero args is fine -> 0

	// Spread an existing slice into a variadic function with ...
	nums := []int{5, 5, 5, 5}
	fmt.Println("sum(nums...)     =", sum(nums...))
}
