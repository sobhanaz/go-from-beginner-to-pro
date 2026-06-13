// Session 07 — bytes vs runes, and building strings efficiently.
// Run:  go run examples/session07/runes/runes.go
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Reverse correctly by operating on RUNES, not bytes.
func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // multiple assignment swap
	}
	return string(r)
}

func main() {
	word := "héllo"

	// Ranging over a string yields RUNES with their starting byte index.
	for i, r := range word {
		fmt.Printf("byte %d: %c (rune code %d)\n", i, r, r)
	}

	// Byte length vs character length.
	fmt.Println("len (bytes):     ", len(word))                  // 6
	fmt.Println("rune count:      ", len([]rune(word)))          // 5
	fmt.Println("utf8.RuneCount:  ", utf8.RuneCountInString(word)) // 5

	// Reversing: byte-reverse would corrupt multi-byte chars; rune-reverse works.
	fmt.Println("reverse:", reverse(word)) // olléh

	// Efficient concatenation with strings.Builder.
	var b strings.Builder
	for i := 0; i < 5; i++ {
		b.WriteString("go ")
	}
	fmt.Printf("builder result: %q\n", b.String())
}
