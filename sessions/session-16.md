> 🌐 **Language / زبان:** English (this file) · [فارسی](session-16.fa.md)

# Session 16 — Testing 🧪

**Goal (1 hour):** Learn Go's built-in testing — no third-party framework needed.
You'll write unit tests, the idiomatic **table-driven** style, subtests, and
benchmarks, and measure **coverage**. Testing is a skill that loudly signals
"professional developer" to employers, and every serious Go codebase is tested.

> **Recap from Session 15:** you can move data via JSON and files. Now you learn
> to *prove your code works* — and keep it working as it changes.

---

## 1. Your first test (15 min)

Go testing is part of the toolchain. The rules:

- Test files end in **`_test.go`** and live **next to** the code they test.
- Test functions start with **`Test`**, are capitalized, and take **`*testing.T`**.
- You run them with **`go test`**.

Given a function in `calc.go`:

```go
func Add(a, b int) int { return a + b }
```

The test in `calc_test.go`:

```go
package calc

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

The pattern: compute `got`, compare to `want`, and call `t.Errorf` (with a clear
message) when they differ. There are no `assert` keywords — you use plain `if`.

| Method | Effect |
|--------|--------|
| `t.Errorf(...)` | mark the test failed, **keep running** |
| `t.Fatalf(...)` | mark failed and **stop this test immediately** (use when continuing is pointless, e.g. after an unexpected error) |

Run it:

```bash
go test ./examples/session16/calc/
go test -v ./examples/session16/calc/   # verbose: shows each test
```

---

## 2. Table-driven tests — the Go idiom (20 min)

Instead of writing a separate function per case, Go developers put cases in a
**slice of structs** and loop over them. This is the single most common testing
pattern in real Go code:

```go
func TestIsEven(t *testing.T) {
    cases := []struct {
        name string
        in   int
        want bool
    }{
        {"zero is even", 0, true},
        {"three is odd", 3, false},
        {"negative even", -4, true},
    }

    for _, c := range cases {
        t.Run(c.name, func(t *testing.T) {   // a named SUBTEST per case
            if got := IsEven(c.in); got != c.want {
                t.Errorf("IsEven(%d) = %v; want %v", c.in, got, c.want)
            }
        })
    }
}
```

Why this is loved:
- **Add a case = add one line.** Easy to extend.
- **`t.Run` makes subtests** — failures name the exact case (`TestIsEven/three_is_odd`).
- It reads like a **specification** of the function's behavior.

> 🔑 This table-driven + `t.Run` style is what interviewers expect to see. Make it
> your default.

### Testing error returns

For the `value, err` pattern, test both the happy and the failing path — and use
`errors.Is` to check for a specific error (Session 11):

```go
_, err := Divide(10, 0)
if !errors.Is(err, ErrDivideByZero) {
    t.Errorf("expected ErrDivideByZero, got %v", err)
}
```

Run [`examples/session16/calc/`](../examples/session16/calc/) and
[`examples/session16/stringutil/`](../examples/session16/stringutil/):

```bash
go test -v ./examples/session16/...
```

---

## 3. Coverage — how much is tested? (10 min)

Go measures what fraction of your statements your tests exercise:

```bash
go test -cover ./examples/session16/...
# ok  golearn/examples/session16/calc  coverage: 100.0% of statements
```

For a visual report, generate and open an HTML view:

```bash
go test -coverprofile=cover.out ./examples/session16/calc/
go tool cover -html=cover.out      # opens a browser highlighting tested lines
```

> 💡 **Coverage is a guide, not a goal.** 100% coverage doesn't mean bug-free, and
> chasing 100% everywhere wastes time. Aim to cover the important logic and the
> tricky edge cases. Recruiters love seeing *meaningful* tests, not just a number.

---

## 4. Benchmarks (10 min)

Benchmarks measure performance. They start with **`Benchmark`**, take
**`*testing.B`**, and loop **`b.N`** times (Go picks `b.N` to get a stable measurement):

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
```

```bash
go test -bench=. ./examples/session16/calc/
go test -bench=. -benchmem ./examples/session16/calc/   # also report allocations
```

The output shows nanoseconds per operation (`ns/op`) and, with `-benchmem`, bytes
and allocations per op — invaluable when optimizing.

> 📦 **Helpful testing extras:** `go test ./...` runs every test in the project.
> `t.Helper()` marks a function as a test helper (so failures point at the caller).
> For larger projects the community uses `testify` for nicer assertions, but the
> standard library is enough — and interviewers respect that you know the built-in way.

---

## 🎯 Exercises (do these before Session 17!)

Create a package `examples/session16/practice/` with code **and** a `_test.go`:

1. **Test a function:** Write `Max(a, b int) int` and a `TestMax` covering a>b,
   b>a, and a==b.
2. **Table-driven:** Write `Clamp(n, lo, hi int) int` (clamps n into [lo,hi]) and
   a table-driven test with at least 5 cases, each a named subtest.
3. **Error path:** Write `ParsePositive(s string) (int, error)` that parses an int
   and errors if it's ≤ 0. Test the success case and both failure cases
   (non-numeric, and zero/negative).
4. **Run coverage:** Run `go test -cover` on your package and try to reach high,
   meaningful coverage.
5. **Benchmark:** Write a `Fib(n int) int` (recursive Fibonacci) and a
   `BenchmarkFib`. Run it and read the `ns/op`.

---

## ✅ Session 16 Checklist

- [ ] I know test files end in `_test.go` and live beside the code
- [ ] I can write a `TestX(t *testing.T)` with got/want and `t.Errorf`
- [ ] I know the difference between `t.Errorf` and `t.Fatalf`
- [ ] I can write a table-driven test with `t.Run` subtests
- [ ] I test both the value and the error in `value, err` functions
- [ ] I can measure coverage with `go test -cover`
- [ ] I can write and run a `Benchmark` with `-bench`
- [ ] I completed all 5 exercises

**Previous:** [← Session 15](session-15.md) · **Next:** [Session 17 — HTTP Servers →](session-17.md)
