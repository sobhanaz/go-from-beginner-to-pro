// Session 12 — channels: sending, receiving, buffering, closing, ranging.
// Run:  go run examples/session12/channels/channels.go
package main

import "fmt"

// produce sends 1..5 then closes the channel to signal "done".
func produce(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // SEND into the channel
	}
	close(ch) // only the SENDER closes
}

func main() {
	// 1. Unbuffered channel synchronizes two goroutines.
	greeting := make(chan string)
	go func() {
		greeting <- "result from goroutine" // blocks until main receives
	}()
	msg := <-greeting // blocks until the goroutine sends -> this is our "wait"
	fmt.Println("received:", msg)

	// 2. Buffered channel holds values without an immediate receiver.
	buf := make(chan int, 3)
	buf <- 10
	buf <- 20
	buf <- 30 // doesn't block until the buffer (cap 3) is full
	close(buf)
	fmt.Println("buffered:", <-buf, <-buf, <-buf)

	// 3. Range over a channel reads until it is closed.
	nums := make(chan int)
	go produce(nums)
	sum := 0
	for n := range nums {
		sum += n
	}
	fmt.Println("sum of 1..5 from channel:", sum)

	// 4. comma-ok detects a closed+drained channel.
	done := make(chan int)
	close(done)
	if v, ok := <-done; !ok {
		fmt.Printf("channel closed; got zero value %d, ok=%t\n", v, ok)
	}
}
