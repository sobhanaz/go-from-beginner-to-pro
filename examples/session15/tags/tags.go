// Session 15 — struct tags control JSON field names, omission, and skipping.
// Run:  go run examples/session15/tags/tags.go
package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Discount float64 `json:"discount,omitempty"` // omitted from JSON when zero
	Internal string  `json:"-"`                   // NEVER serialized to JSON
}

func main() {
	// Discount is 0 (zero value) and Internal is set but tagged "-".
	p1 := Product{ID: 1, Title: "Keyboard", Price: 49.99, Internal: "secret-note"}
	out1, _ := json.MarshalIndent(p1, "", "  ")
	fmt.Printf("no discount (omitempty hides it; Internal hidden):\n%s\n\n", out1)

	// Now discount is non-zero, so it appears.
	p2 := Product{ID: 2, Title: "Mouse", Price: 19.99, Discount: 5.0, Internal: "x"}
	out2, _ := json.MarshalIndent(p2, "", "  ")
	fmt.Printf("with discount:\n%s\n", out2)

	// Unmarshalling: JSON keys map back via the tags. Unknown keys are ignored;
	// missing keys keep their zero value.
	input := []byte(`{"id": 3, "title": "Cable", "price": 7.5, "extra": "ignored"}`)
	var p3 Product
	_ = json.Unmarshal(input, &p3)
	fmt.Printf("\nparsed (extra ignored, discount=0): %+v\n", p3)
}
