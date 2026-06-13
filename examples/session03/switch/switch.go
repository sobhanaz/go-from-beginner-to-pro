// Session 03 — the many faces of switch.
// Run:  go run examples/session03/switch/switch.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. Classic switch on a value. No `break` needed.
	//    One case can match multiple values.
	day := "Sat"
	switch day {
	case "Sat", "Sun":
		fmt.Println("Weekend 🎉")
	case "Mon":
		fmt.Println("Monday...")
	default:
		fmt.Println("A weekday")
	}

	// 2. Condition-less switch: cleanest way to do range/threshold checks.
	score := 83
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	// 3. switch with a short statement (initialize, then test).
	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	// 4. fallthrough (rare) forces continuing into the next case.
	switch n := 1; n {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("...and we fell through to two")
	case 3:
		fmt.Println("three (not reached)")
	}
}
