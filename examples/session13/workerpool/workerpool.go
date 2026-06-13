// Session 13 — the worker pool pattern (goroutines + channels + WaitGroup).
// Run:  go run -race examples/session13/workerpool/workerpool.go
package main

import (
	"fmt"
	"sort"
	"sync"
)

// Each worker pulls jobs until the jobs channel closes, pushing results out.
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * j // the "work": square the number
	}
}

func main() {
	const numWorkers = 3
	jobs := make(chan int, 9)
	results := make(chan int, 9)
	var wg sync.WaitGroup

	// Start a fixed number of workers.
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs, then close so workers' `range` loops end.
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Close results once all workers are done (in a goroutine so we can range below).
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results (order is non-deterministic, so sort for stable output).
	var collected []int
	for r := range results {
		collected = append(collected, r)
	}
	sort.Ints(collected)
	fmt.Println("squares:", collected)
}
