// Session 11 — panic and recover (the emergency exit).
// Run:  go run examples/session11/panic/panic.go
package main

import "fmt"

// safeRun runs a task and recovers if it panics, so the program survives.
// This is the shape of "recovery middleware" used in web servers.
func safeRun(name string, task func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("[%s] recovered from panic: %v\n", name, r)
		}
	}()
	task()
	fmt.Printf("[%s] completed normally\n", name)
}

// Turn a divide-by-zero panic into a returned error (named returns let the
// deferred function set err).
func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	result = a / b // panics if b == 0
	return result, nil
}

func main() {
	safeRun("task-1", func() {
		fmt.Println("doing safe work")
	})

	safeRun("task-2", func() {
		panic("something exploded")
	})

	fmt.Println("main is still running after the panic!")

	if r, err := safeDivide(10, 2); err == nil {
		fmt.Println("10 / 2 =", r)
	}
	if _, err := safeDivide(10, 0); err != nil {
		fmt.Println("10 / 0 ->", err)
	}
}
