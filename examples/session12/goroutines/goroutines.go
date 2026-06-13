// Session 12 — starting goroutines, and why main must wait.
// Run:  go run examples/session12/goroutines/goroutines.go
package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(50 * time.Millisecond) // simulate work
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// Launch several goroutines — they run concurrently.
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	// PROBLEM: main would exit immediately and kill the goroutines.
	// Here we use a crude Sleep just to demonstrate. (Session 13 shows the
	// proper way with sync.WaitGroup; Session 12 also shows channels.)
	fmt.Println("main launched 3 workers, waiting...")
	time.Sleep(200 * time.Millisecond)
	fmt.Println("main finished")
}
