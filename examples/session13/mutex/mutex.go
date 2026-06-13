// Session 13 — sync.Mutex protects shared data from races.
// Run:  go run -race examples/session13/mutex/mutex.go
//
// Experiment: comment out the Lock/Unlock lines and rerun with -race —
// the race detector will report a data race and the total will be wrong.
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()         // only one goroutine in the critical section at a time
	defer c.mu.Unlock() // released when Inc returns
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	c := &Counter{}
	var wg sync.WaitGroup

	const n = 1000
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()
	fmt.Printf("final count: %d (expected %d)\n", c.Value(), n)
}
