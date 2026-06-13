// Session 03 — iota for auto-numbered constants (enums).
// Run:  go run examples/session03/iota/iota.go
package main

import "fmt"

// iota starts at 0 and increments by 1 on each line of the const block.
const (
	Sunday = iota // 0
	Monday        // 1
	Tuesday       // 2
	Wednesday     // 3
)

// A named type makes an enum that is self-documenting and type-safe.
type Status int

const (
	StatusPending Status = iota // 0
	StatusActive                // 1
	StatusClosed                // 2
)

// You can attach a String() method so the enum prints nicely
// (a preview of "methods", coming in Session 08).
func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusActive:
		return "active"
	case StatusClosed:
		return "closed"
	default:
		return "unknown"
	}
}

func main() {
	fmt.Println("Days:", Sunday, Monday, Tuesday, Wednesday)

	task := StatusActive
	fmt.Printf("Task status value: %d, name: %v\n", task, task)

	// Because of the String() method, printing shows the friendly name.
	fmt.Println("All statuses:", StatusPending, StatusActive, StatusClosed)
}
