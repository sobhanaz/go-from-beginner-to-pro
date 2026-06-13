// Session 07 — maps: key/value lookup.
// Run:  go run examples/session07/maps/maps.go
package main

import "fmt"

func main() {
	// Map literal: map[KeyType]ValueType
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}

	// Read, add, update, delete.
	fmt.Println("Alice:", ages["Alice"])
	ages["Carol"] = 28 // add
	ages["Alice"] = 31 // update
	delete(ages, "Bob")
	fmt.Println("after edits:", ages, "len:", len(ages))

	// Reading a MISSING key returns the zero value (0 here) — ambiguous!
	fmt.Println("Dave (missing):", ages["Dave"]) // 0

	// comma-ok idiom: tells you whether the key actually exists.
	if age, ok := ages["Dave"]; ok {
		fmt.Println("Dave is", age)
	} else {
		fmt.Println("No Dave in the map")
	}

	// Iterating — ORDER IS RANDOM by design.
	for name, age := range ages {
		fmt.Printf("%s is %d\n", name, age)
	}

	// A small word-count, the canonical map use-case.
	counts := map[string]int{}
	words := []string{"go", "is", "fun", "go", "go"}
	for _, w := range words {
		counts[w]++ // missing key starts at 0, then increments
	}
	fmt.Println("word counts:", counts)
}
