// Session 09 — pointer basics: & and *.
// Run:  go run examples/session09/basics/basics.go
package main

import "fmt"

func main() {
	x := 10
	p := &x // p is a *int holding the address of x

	fmt.Println("x  =", x)  // the value
	fmt.Println("p  =", p)  // an address like 0xc0000...
	fmt.Println("*p =", *p) // dereference: the value at that address

	// Writing through the pointer changes the original variable.
	*p = 20
	fmt.Println("after *p = 20, x =", x) // 20

	// The zero value of a pointer is nil (points to nothing).
	var np *int
	fmt.Println("nil pointer:", np) // <nil>
	if np == nil {
		fmt.Println("np is nil — dereferencing it would panic")
	}
}
