// Session 10 — the empty interface (any), type assertions, type switch.
// Run:  go run examples/session10/empty/empty.go
package main

import "fmt"

// describe accepts ANY value and reports its type with a type switch.
func describe(i any) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %d (doubled: %d)\n", v, v*2)
	case string:
		fmt.Printf("string: %q (length %d)\n", v, len(v))
	case bool:
		fmt.Printf("bool: %t\n", v)
	case float64:
		fmt.Printf("float64: %.2f\n", v)
	default:
		fmt.Printf("unknown type %T: %v\n", v, v)
	}
}

func main() {
	// `any` (alias for interface{}) can hold any value.
	var x any
	x = 42
	fmt.Println("x holds:", x)
	x = "hello"

	// Type assertion (safe comma-ok form).
	if s, ok := x.(string); ok {
		fmt.Println("recovered string:", s)
	}

	// A wrong assertion with comma-ok does NOT panic — ok is just false.
	if n, ok := x.(int); ok {
		fmt.Println("it's an int:", n)
	} else {
		fmt.Println("x is not an int")
	}

	// Type switch handles many possibilities cleanly.
	describe(7)
	describe("golang")
	describe(true)
	describe(3.14159)
	describe([]int{1, 2, 3})
}
