// Session 02 — type conversion and the integer-division trap.
// Run:  go run examples/session02/conversion/conversion.go
package main

import "fmt"

func main() {
	// Go will NOT mix types automatically. You must convert explicitly.
	var x int = 5
	var y float64 = 2.5

	// z := x + y          // <-- would be a COMPILE ERROR: mismatched types
	z := float64(x) + y     // convert x to float64 first
	fmt.Printf("float64(x) + y = %v\n", z) // 7.5

	// The classic integer-division trap:
	a := 10
	b := 3
	fmt.Println("10 / 3 (int)   =", a/b)                       // 3  (remainder thrown away)
	fmt.Println("10 / 3 (float) =", float64(a)/float64(b))     // 3.3333333333333335

	// Truncation when converting float -> int (NOT rounding).
	height := 1.81
	fmt.Printf("int(1.81) = %d (decimal part dropped)\n", int(height)) // 1

	// Converting between integer sizes is also explicit.
	var big int64 = 1000
	var small int32 = int32(big)
	fmt.Printf("int64 %d -> int32 %d\n", big, small)
}
