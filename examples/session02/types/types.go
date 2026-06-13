// Session 02 — the basic types and their zero values.
// Run:  go run examples/session02/types/types.go
package main

import "fmt"

// Constants live fine at package level (top of file, outside functions).
const AppName = "TaskFlow"
const Pi = 3.14159

func main() {
	// The four types you'll use 95% of the time.
	var text string = "hello"
	var whole int = 42
	var decimal float64 = 3.14
	var flag bool = true

	fmt.Printf("string  -> value:%q  type:%T\n", text, text)
	fmt.Printf("int     -> value:%d  type:%T\n", whole, whole)
	fmt.Printf("float64 -> value:%v  type:%T\n", decimal, decimal)
	fmt.Printf("bool    -> value:%t  type:%T\n", flag, flag)

	// Underscores improve readability of big numbers (ignored by Go).
	population := 9_500_000
	fmt.Println("population:", population)

	// Zero values — declared with no assignment.
	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool
	fmt.Printf("zero values -> int:%d float:%v string:%q bool:%t\n",
		zeroInt, zeroFloat, zeroString, zeroBool)

	fmt.Println("App:", AppName, "Pi:", Pi)
}
