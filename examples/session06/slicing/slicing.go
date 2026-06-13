// Session 06 — slicing syntax and the shared-memory gotcha.
// Run:  go run examples/session06/slicing/slicing.go
package main

import "fmt"

func main() {
	nums := []int{10, 20, 30, 40, 50}

	// s[low:high] — low included, high EXCLUDED.
	fmt.Println("nums[1:3] =", nums[1:3]) // [20 30]
	fmt.Println("nums[:2]  =", nums[:2])  // [10 20]
	fmt.Println("nums[3:]  =", nums[3:])  // [40 50]
	fmt.Println("nums[:]   =", nums[:])   // whole slice

	// THE GOTCHA: a sub-slice shares the original's underlying array.
	original := []int{1, 2, 3, 4, 5}
	part := original[1:3] // [2 3], shares memory
	part[0] = 99
	fmt.Println("after part[0]=99 -> original:", original) // [1 99 3 4 5]

	// copy() makes an independent duplicate.
	src := []int{1, 2, 3, 4, 5}
	view := src[1:3]
	dst := make([]int, len(view))
	copy(dst, view) // dst is now separate
	dst[0] = 99
	fmt.Println("with copy -> src:", src, "dst:", dst) // src unchanged
}
