// Session 16 — table-driven tests for stringutil.
// Run:  go test -v ./examples/session16/stringutil/
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"héllo", "olléh"}, // multi-byte rune handled correctly
	}
	for _, c := range cases {
		t.Run(c.in, func(t *testing.T) {
			if got := Reverse(c.in); got != c.want {
				t.Errorf("Reverse(%q) = %q; want %q", c.in, got, c.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want bool
	}{
		{"simple palindrome", "racecar", true},
		{"not a palindrome", "hello", false},
		{"with spaces and case", "Never odd or even", true},
		{"empty is palindrome", "", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := IsPalindrome(c.in); got != c.want {
				t.Errorf("IsPalindrome(%q) = %v; want %v", c.in, got, c.want)
			}
		})
	}
}
