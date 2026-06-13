// Session 11 — sentinel errors (errors.Is) and custom error types (errors.As).
// Run:  go run examples/session11/custom/custom.go
package main

import (
	"errors"
	"fmt"
)

// Sentinel error: a predefined, comparable value. Convention: name starts "Err".
var ErrNotFound = errors.New("not found")

func findUser(id int) (string, error) {
	if id == 0 {
		// Wrap it to add context; errors.Is still finds it through the wrap.
		return "", fmt.Errorf("findUser(%d): %w", id, ErrNotFound)
	}
	return "Sobhan", nil
}

// Custom error struct that carries extra data.
type ValidationError struct {
	Field string
	Msg   string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Msg: "must be non-negative"}
	}
	return nil
}

func main() {
	// errors.Is: detect a SPECIFIC error value, even when wrapped.
	_, err := findUser(0)
	if errors.Is(err, ErrNotFound) {
		fmt.Println("detected ErrNotFound (message:", err, ")")
	}

	// errors.As: detect an error of a specific TYPE and extract its fields.
	err = validateAge(-5)
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("validation problem -> field:%q msg:%q\n", ve.Field, ve.Msg)
	}
}
