// Session 15 — encoding/json: Marshal (struct->JSON) and Unmarshal (JSON->struct).
// Run:  go run examples/session15/jsonbasics/jsonbasics.go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// --- Marshal: Go value -> JSON bytes ---
	u := User{ID: 1, Name: "Sobhan", Email: "sobhan@example.com"}

	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}
	fmt.Println("compact JSON:", string(data))

	// MarshalIndent produces pretty, human-readable JSON.
	pretty, _ := json.MarshalIndent(u, "", "  ")
	fmt.Printf("pretty JSON:\n%s\n", pretty)

	// --- Unmarshal: JSON bytes -> Go value ---
	input := []byte(`{"id": 2, "name": "Ali", "email": "ali@example.com"}`)
	var parsed User
	if err := json.Unmarshal(input, &parsed); err != nil { // note the &
		fmt.Println("unmarshal error:", err)
		return
	}
	fmt.Printf("parsed struct: %+v\n", parsed)

	// Slices marshal to JSON arrays.
	users := []User{u, parsed}
	arr, _ := json.Marshal(users)
	fmt.Println("array JSON:", string(arr))
}
