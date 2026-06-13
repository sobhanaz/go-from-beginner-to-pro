// Session 08 — defining and using structs.
// Run:  go run examples/session08/structs/structs.go
package main

import "fmt"

// A custom type that groups related fields.
type User struct {
	ID     int
	Name   string
	Email  string
	Active bool
}

func main() {
	// 1. Named fields — clearest, order-independent.
	u1 := User{
		ID:     1,
		Name:   "Sobhan",
		Email:  "sobhan@example.com",
		Active: true,
	}

	// 2. Positional — must match field order.
	u2 := User{2, "Ali", "ali@example.com", false}

	// 3. Zero value — every field gets its type's zero value.
	var u3 User

	// %v prints values; %+v also prints field NAMES (very handy for debugging).
	fmt.Printf("u1: %+v\n", u1)
	fmt.Printf("u2: %v\n", u2)
	fmt.Printf("u3 (zero): %+v\n", u3)

	// Access and modify fields.
	fmt.Println("u1.Name:", u1.Name)
	u1.Name = "Sobhan A."
	u1.Active = false
	fmt.Printf("u1 after edit: %+v\n", u1)
}
