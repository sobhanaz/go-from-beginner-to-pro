// Session 04 — FizzBuzz, the classic interview question, done idiomatically.
// Run:  go run examples/session04/fizzbuzz/fizzbuzz.go
package main

import "fmt"

func fizzbuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0: // divisible by BOTH 3 and 5 -> check first!
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		// case i%2 == 0:
		// 	fmt.Println("Even")
		// case i%2 != 0:
		// 	fmt.Println("Odd")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	fizzbuzz(30)
}
