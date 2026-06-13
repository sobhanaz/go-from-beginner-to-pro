// Session 04 — every shape of the for loop.
// Run:  go run examples/session04/loops/loops.go
package main

import "fmt"

func main() {
	// Shape 1: classic three-part loop.
	fmt.Print("count up:   ")
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// Shape 2: condition-only ("while" style).
	n := 1
	for n < 100 {
		n = n * 2
	}
	fmt.Println("doubled past 100:", n) // 128

	// Shape 3: infinite loop with break.
	i := 0
	for {
		if i >= 3 {
			break
		}
		fmt.Println("tick", i)
		i++
	}

	// break and continue together.
	fmt.Print("skip 3, stop at 6: ")
	for j := 0; j < 10; j++ {
		if j == 3 {
			continue // skip this iteration
		}
		if j == 6 {
			break // exit the loop
		}
		fmt.Print(j, " ")
	}
	fmt.Println()

	// Shape 4: for range over a slice (preview of Session 06).
	names := []string{"Ann", "Bob", "Cy"}
	for index, name := range names {
		fmt.Printf("%d -> %s\n", index, name)
	}

	// Use _ to ignore the index when you only want the value.
	total := 0
	nums := []int{10, 20, 30}
	for _, v := range nums {
		total += v
	}
	fmt.Println("sum of nums:", total)
}
