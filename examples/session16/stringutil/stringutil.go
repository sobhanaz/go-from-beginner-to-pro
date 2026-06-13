// Session 16 — a second small package to practice testing on.
package stringutil

import "strings"

// Reverse returns s reversed, correctly handling multi-byte (rune) characters.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// IsPalindrome reports whether s reads the same forwards and backwards,
// ignoring case and spaces.
func IsPalindrome(s string) bool {
	clean := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return clean == Reverse(clean)
}
