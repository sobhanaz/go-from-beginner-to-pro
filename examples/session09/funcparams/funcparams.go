// Session 09 — why functions need pointers to modify their arguments.
// Run:  go run examples/session09/funcparams/funcparams.go
package main

import "fmt"

// Takes a COPY — changes are lost.
func tryToDouble(n int) {
	n = n * 2
}

// Takes a POINTER — modifies the caller's real value.
func double(n *int) {
	*n = *n * 2
}

func main() {
	x := 5

	tryToDouble(x)
	fmt.Println("after tryToDouble:", x) // still 5

	double(&x) // pass the ADDRESS of x
	fmt.Println("after double:     ", x) // 10
}
