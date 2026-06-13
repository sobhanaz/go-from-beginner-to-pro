// Session 06 — slices: the dynamic, growable list.
// Run:  go run examples/session06/slices/slices.go
package main

import "fmt"

func main() {
	// Slice literal (no size in the brackets).
	fruits := []string{"apple", "banana"}
	fmt.Println("fruits:", fruits, "len:", len(fruits))

	// append returns a NEW slice — always assign it back.
	fruits = append(fruits, "cherry")
	fruits = append(fruits, "date", "elderberry") // several at once
	fmt.Println("after append:", fruits)

	// Spread one slice into append with ...
	more := []string{"fig", "grape"}
	fruits = append(fruits, more...)
	fmt.Println("after spread:", fruits)

	// make([]T, length, capacity)
	scores := make([]int, 0, 5)
	fmt.Printf("scores -> len:%d cap:%d\n", len(scores), cap(scores))
	for i := 1; i <= 6; i++ {
		scores = append(scores, i*10)
	}
	// Capacity grew automatically once we passed 5.
	fmt.Printf("scores:%v len:%d cap:%d\n", scores, len(scores), cap(scores))

	// Building a new slice from an old one ("map" in other languages).
	prices := []float64{10, 20, 30}
	withTax := make([]float64, 0, len(prices))
	for _, p := range prices {
		withTax = append(withTax, p*1.09)
	}
	fmt.Println("withTax:", withTax)
}
