// Session 13 — sync.WaitGroup: wait for many goroutines correctly.
// Run:  go run -race examples/session13/waitgroup/waitgroup.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1) // BEFORE launching: "one more goroutine to wait for"
		go func(id int) {
			defer wg.Done() // "this one is finished" when it returns
			fmt.Printf("worker %d doing work\n", id)
		}(i) // pass i as an argument to avoid loop-variable capture bugs
	}

	wg.Wait() // block until the counter returns to 0
	fmt.Println("all workers finished")
}
