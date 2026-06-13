// Session 16 — tests for the calc package.
// Test files end in _test.go and live next to the code they test.
// Run:  go test ./examples/session16/calc/
//       go test -v ./examples/session16/calc/   (verbose)
//       go test -cover ./examples/session16/calc/  (coverage)
package calc

import (
	"errors"
	"testing"
)

// A basic test: function name starts with Test and takes *testing.T.
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		// t.Errorf reports a failure but keeps running.
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
}

// TABLE-DRIVEN test: the idiomatic Go style. One test, many cases.
func TestIsEven(t *testing.T) {
	cases := []struct {
		name string
		in   int
		want bool
	}{
		{"zero is even", 0, true},
		{"two is even", 2, true},
		{"three is odd", 3, false},
		{"negative even", -4, true},
		{"negative odd", -7, false},
	}

	for _, c := range cases {
		// t.Run creates a named SUBTEST so failures point to the exact case.
		t.Run(c.name, func(t *testing.T) {
			if got := IsEven(c.in); got != c.want {
				t.Errorf("IsEven(%d) = %v; want %v", c.in, got, c.want)
			}
		})
	}
}

// Testing the value AND the error return.
func TestDivide(t *testing.T) {
	t.Run("normal division", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("unexpected error: %v", err) // Fatalf stops THIS subtest
		}
		if got != 5 {
			t.Errorf("Divide(10, 2) = %v; want 5", got)
		}
	})

	t.Run("divide by zero returns error", func(t *testing.T) {
		_, err := Divide(10, 0)
		if !errors.Is(err, ErrDivideByZero) {
			t.Errorf("expected ErrDivideByZero, got %v", err)
		}
	})
}

// A benchmark: name starts with Benchmark, takes *testing.B, loops b.N times.
// Run:  go test -bench=. ./examples/session16/calc/
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(100, 200)
	}
}
