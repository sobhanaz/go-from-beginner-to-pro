> 🌐 **Language / زبان:** English (this file) · [فارسی](session-05.fa.md)

# Session 05 — Functions Deep Dive 🛠️

**Goal (1 hour):** Go beyond basic functions and learn the features that make Go
*feel* like Go: **multiple return values** (and the `value, err` pattern that's
everywhere), **variadic** functions, **closures**, and **`defer`**. Master these
and you'll read real Go code with ease.

> **Recap from Session 04:** you can loop with `for` and write/call functions
> that take parameters and return a value. Now we supercharge functions.

---

## 1. Multiple return values (15 min) — *the* Go signature

Unlike most languages, a Go function can return **more than one value**. This is
not a rare trick — it's the heart of idiomatic Go.

```go
func divide(a, b int) (int, int) {   // returns quotient AND remainder
    return a / b, a % b
}

q, r := divide(17, 5)   // q=3, r=2
```

### The `value, err` pattern (memorize this!)

The single most important convention in Go: functions that can fail return
**two** things — the result *and* an error. If the error is `nil`, it worked.

```go
func sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, errors.New("cannot sqrt a negative number")
    }
    return math.Sqrt(x), nil   // nil error = success
}

result, err := sqrt(16)
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println("result:", result)
```

> 🔑 **You will write `if err != nil { ... }` thousands of times** in your Go
> career. Go has no exceptions for normal errors — instead, errors are *values*
> you return and check. We go deep on errors in Session 11; here just learn the shape.

When you only care about some returns, discard the rest with `_`:

```go
_, remainder := divide(17, 5)   // ignore the quotient
```

### Named return values

You can name your returns. They act like pre-declared variables, and a bare
`return` sends their current values back:

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return   // "naked return" — sends x and y
}
```

Use named returns sparingly — they're nice for short functions and for
documenting what each return *means*, but can hurt readability in long ones.

Run [`examples/session05/multireturn/multireturn.go`](../examples/session05/multireturn/multireturn.go).

---

## 2. Variadic functions (10 min) — accept any number of args

A `...` before the type means "zero or more of these." Inside the function, the
parameter is a **slice**.

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

sum(1, 2, 3)        // 6
sum(10, 20, 30, 40) // 100
sum()               // 0
```

`fmt.Println` itself is variadic — that's why you can pass it any number of args.

To pass an existing slice into a variadic function, "spread" it with `...`:

```go
nums := []int{1, 2, 3}
sum(nums...)   // spreads the slice into individual args
```

Run [`examples/session05/variadic/variadic.go`](../examples/session05/variadic/variadic.go).

---

## 3. Functions are values; closures (15 min)

In Go, functions are **first-class values** — you can store them in variables,
pass them as arguments, and return them from other functions.

```go
add := func(a, b int) int {   // anonymous function stored in a variable
    return a + b
}
fmt.Println(add(2, 3))   // 5
```

A **closure** is a function that "closes over" (remembers) variables from the
scope where it was created. This is powerful:

```go
func counter() func() int {
    count := 0
    return func() int {   // this inner func remembers `count`
        count++
        return count
    }
}

next := counter()
fmt.Println(next()) // 1
fmt.Println(next()) // 2
fmt.Println(next()) // 3
```

Each call to `counter()` makes a fresh, independent `count`. Closures are how Go
does stateful behavior without classes. You'll see them in middleware, handlers,
and callbacks throughout your career.

Run [`examples/session05/closures/closures.go`](../examples/session05/closures/closures.go).

---

## 4. `defer` — cleanup that always runs (15 min)

`defer` schedules a function call to run **when the surrounding function
returns** — no matter how it returns (normal return, or even a panic). It's Go's
tool for cleanup.

```go
func readFile() {
    file := open("data.txt")
    defer file.Close()   // guaranteed to run when readFile returns

    // ... do work with file ...
    // even if we return early or panic, file.Close() still runs
}
```

> 🔑 **Why this matters:** you write the cleanup *right next to* the thing that
> needs cleaning up, so you can't forget it. You'll use `defer` constantly to
> close files, unlock mutexes, and close database connections.

**Two rules to remember:**

1. **LIFO order:** multiple `defer`s run in reverse (Last In, First Out).
   ```go
   defer fmt.Println("1")
   defer fmt.Println("2")
   defer fmt.Println("3")
   // prints: 3, 2, 1
   ```
2. **Arguments are evaluated immediately**, but the call happens at the end:
   ```go
   x := 10
   defer fmt.Println("deferred x:", x) // captures 10 NOW
   x = 20
   // at return, prints "deferred x: 10"
   ```

Run [`examples/session05/defer/defer.go`](../examples/session05/defer/defer.go).

---

## 🎯 Exercises (do these before Session 06!)

Create `examples/session05/practice/practice.go`:

1. **min & max:** Write `func minMax(numbers ...int) (int, int)` that returns the
   smallest and largest of any number of ints. Call it and print both.
2. **Safe divide:** Write `func safeDivide(a, b float64) (float64, error)` that
   returns an error when `b == 0` (use `errors.New`). Test it with `10/2` and
   `10/0`, handling the error with `if err != nil`.
3. **Average (variadic):** Write `func average(nums ...float64) float64`. Bonus:
   handle the empty case so you don't divide by zero.
4. **Closure multiplier:** Write `func multiplier(factor int) func(int) int` that
   returns a function multiplying its input by `factor`. Make a `double` and a
   `triple` from it.
5. **defer order:** Write a function with three `defer fmt.Println(...)` calls and
   predict the output order before running it.

---

## ✅ Session 05 Checklist

- [ ] I can return multiple values from a function
- [ ] I understand and can write the `value, err :=` + `if err != nil` pattern
- [ ] I know what named returns and a "naked return" are
- [ ] I can write and call a variadic `...` function, and spread a slice with `...`
- [ ] I can store a function in a variable and write a closure
- [ ] I can explain what `defer` does and its LIFO order
- [ ] I completed all 5 exercises

**Previous:** [← Session 04](session-04.md) · **Next:** [Session 06 — Arrays & Slices →](session-06.md)

---

🎉 **Milestone:** That's Part 1 (Foundations) complete! You now know Go's core
syntax. Next we tackle data structures — slices, maps, structs — where Go
programs really come alive.
