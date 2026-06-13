// Session 12 — directional channels document and enforce send/receive intent.
// Run:  go run examples/session12/directions/directions.go
package main

import "fmt"

// out is SEND-ONLY: this function may only send into it.
func producer(out chan<- int, count int) {
	for i := 1; i <= count; i++ {
		out <- i * i // send squares
	}
	close(out)
}

// in is RECEIVE-ONLY: this function may only receive from it.
func consumer(in <-chan int, done chan<- bool) {
	for v := range in {
		fmt.Println("got:", v)
	}
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)

	go producer(ch, 5)
	go consumer(ch, done)

	<-done // wait for the consumer to finish
	fmt.Println("pipeline complete")
}
