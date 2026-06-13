// Session 02 — declaring variables the different ways.
// Run:  go run examples/session02/variables/variables.go
package main

import "fmt"

func main() {
	// --- Way 1: var with explicit type ---
	var name string = "Sobhan"
	var age int = 25

	// var with type inference (Go figures out the type)
	var city = "Tehran"

	// var with no value -> gets the ZERO value
	var score int     // 0
	var nickname string // ""

	// --- Way 2: short declaration := (most common, function-only) ---
	job := "Go developer in training"
	salaryGoal := 3000

	// Reassigning an existing variable uses = (no colon)
	age = 26

	// Declaring several at once
	x, y := 10, 20

	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("city:", city)
	fmt.Printf("score (zero value): %d\n", score)
	fmt.Printf("nickname (zero value): %q\n", nickname) // %q shows quotes -> ""
	fmt.Println("job:", job)
	fmt.Println("salaryGoal:", salaryGoal)
	fmt.Println("x + y =", x+y)
}
