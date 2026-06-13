// Session 01 — printing and format verbs.
// Run it with:  go run examples/session01/printing/printing.go
package main

import "fmt"

func main() {
	// 1. Println: adds spaces between args and a newline at the end.
	fmt.Println("Hello", "Go", "learner")

	// 2. Print: prints exactly what you give it, no spaces/newline added.
	fmt.Print("no")
	fmt.Print("newline")
	fmt.Print("\n") // we add the newline ourselves

	// 3. Printf: formatted printing with "verbs" (placeholders).
	name := "Sobhan"
	age := 25
	height := 1.81
	learning := true

	fmt.Printf("Name: %s\n", name)      // %s = string
	fmt.Printf("Age: %d\n", age)        // %d = integer
	fmt.Printf("Height: %.2f m\n", height) // %.2f = float with 2 decimals
	fmt.Printf("Learning Go: %t\n", learning) // %t = boolean

	// %v prints ANY value in its default format — your go-to verb.
	fmt.Printf("All together with %%v: %v, %v, %v, %v\n", name, age, height, learning)

	// %T prints the TYPE of a value — super useful while learning.
	fmt.Printf("Types: %T, %T, %T, %T\n", name, age, height, learning)
}
