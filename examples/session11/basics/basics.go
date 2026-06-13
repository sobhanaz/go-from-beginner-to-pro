// Session 11 — error basics: returning, checking, and wrapping.
// Run:  go run examples/session11/basics/basics.go
package main

import (
	"errors"
	"fmt"
	"strconv"
)

// A function that can fail returns an error as its LAST value.
func half(n int) (int, error) {
	if n%2 != 0 {
		return 0, errors.New("number is not even")
	}
	return n / 2, nil // nil error = success
}

// Wrapping with %w preserves the original error while adding context.
func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAge(%q): %w", s, err)
	}
	return age, nil
}

func main() {
	// The golden pattern: check the error immediately.
	if h, err := half(10); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("half of 10 =", h)
	}

	if _, err := half(7); err != nil {
		fmt.Println("half(7) failed:", err)
	}

	// Formatted error with fmt.Errorf.
	if age, err := parseAge("25"); err == nil {
		fmt.Println("parsed age:", age)
	}

	// A wrapped error shows the whole trail.
	if _, err := parseAge("abc"); err != nil {
		fmt.Println("wrapped error:", err)
	}
}
