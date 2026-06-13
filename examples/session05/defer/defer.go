// Session 05 — defer: cleanup that always runs, in LIFO order.
// Run:  go run examples/session05/defer/defer.go
package main

import "fmt"

func main() {
	fmt.Println("start of main")

	// Deferred calls run when main returns, in reverse (LIFO) order.
	defer fmt.Println("deferred 1 (runs last)")
	defer fmt.Println("deferred 2")
	defer fmt.Println("deferred 3 (runs first)")

	// Arguments to a deferred call are evaluated NOW, not at return time.
	x := 10
	defer fmt.Println("deferred x captured at schedule time:", x) // captures 10
	x = 20
	fmt.Println("x is now", x)

	fmt.Println("end of main body")
	// As main returns, deferred calls fire: 3, 2, x=10, 1
}
