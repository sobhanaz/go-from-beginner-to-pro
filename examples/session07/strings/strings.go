// Session 07 — strings are immutable UTF-8 bytes; common operations.
// Run:  go run examples/session07/strings/strings.go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello"
	fmt.Println("len(s):", len(s))     // 5 bytes
	fmt.Println("s[0] (byte):", s[0])  // 72, the byte value of 'H'
	fmt.Printf("s[0] as char: %c\n", s[0]) // H

	// len counts BYTES, not characters.
	fmt.Println(`len("héllo"):`, len("héllo")) // 6, because é is 2 bytes

	// Common strings package helpers.
	fmt.Println(strings.ToUpper("hello"))             // HELLO
	fmt.Println(strings.Contains("hello", "ell"))     // true
	fmt.Println(strings.Split("a,b,c", ","))          // [a b c]
	fmt.Println(strings.ReplaceAll("aaa", "a", "b"))  // bbb
	fmt.Printf("%q\n", strings.TrimSpace("  hi  "))   // "hi"
	fmt.Println(strings.Join([]string{"a", "b"}, "-")) // a-b
	fmt.Println(strings.HasPrefix("golang", "go"))    // true
}
