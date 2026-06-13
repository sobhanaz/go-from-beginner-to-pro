// Session 13 — select (multiple channels, timeouts) and context (cancellation).
// Run:  go run examples/session13/selectctx/selectctx.go
package main

import (
	"context"
	"fmt"
	"time"
)

// doWork returns a channel that delivers a result after `d`.
func doWork(d time.Duration) <-chan string {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(d)
		ch <- fmt.Sprintf("done after %v", d)
	}()
	return ch
}

func main() {
	// 1. select: whichever channel is ready first wins; time.After is a timeout.
	fast := doWork(50 * time.Millisecond)
	select {
	case res := <-fast:
		fmt.Println("fast result:", res)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timed out waiting for fast")
	}

	// 2. select with a timeout that actually fires.
	slow := doWork(500 * time.Millisecond)
	select {
	case res := <-slow:
		fmt.Println("slow result:", res)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("gave up on slow work (timeout)")
	}

	// 3. context with a deadline: cancel work that takes too long.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel() // always release context resources

	work := doWork(300 * time.Millisecond)
	select {
	case res := <-work:
		fmt.Println("ctx work result:", res)
	case <-ctx.Done():
		fmt.Println("context cancelled:", ctx.Err())
	}
}
