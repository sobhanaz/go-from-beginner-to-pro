// Session 03 — if / else if / else and the short-statement form.
// Run:  go run examples/session03/ifelse/ifelse.go
package main

import "fmt"

// A tiny helper so we can show the "if with short statement" pattern.
func getAge() int {
	return 20
}

func main() {
	score := 75

	// Basic if / else if / else. No parentheses; braces required.
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 70 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: Needs improvement")
	}

	// Logical operators combine conditions.
	age := 20
	hasTicket := true
	if age >= 18 && hasTicket {
		fmt.Println("You may enter.")
	}

	// The idiomatic "if with a short statement".
	// The variable `a` only exists inside this if/else block.
	if a := getAge(); a >= 18 {
		fmt.Println("Adult, age is", a)
	} else {
		fmt.Println("Minor, age is", a)
	}

	// Modulo (%) is great for even/odd tests.
	n := 17
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}
