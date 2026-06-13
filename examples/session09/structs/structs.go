// Session 09 — pointers to structs (the common case).
// Run:  go run examples/session09/structs/structs.go
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// Pointer receiver-style function: u.Age auto-dereferences (no (*u).Age needed).
func birthday(u *User) {
	u.Age++
}

func main() {
	user := User{Name: "Sobhan", Age: 25}
	birthday(&user)
	fmt.Printf("after birthday: %+v\n", user) // Age:26

	// Two ways to make a pointer to a fresh struct.
	u1 := &User{Name: "Ali"} // most common
	u2 := new(User)          // new(T) -> *T pointing to a zero value
	u2.Name = "Sara"

	fmt.Printf("u1: %+v\n", *u1)
	fmt.Printf("u2: %+v\n", *u2)
}
