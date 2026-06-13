> 🌐 **Language / زبان:** English (this file) · [فارسی](session-11.fa.md)

# Session 11 — Errors ⚠️

**Goal (1 hour):** Learn how Go handles things going wrong. Go has **no
exceptions** for normal failures — instead, errors are ordinary **values** you
return and check. You'll master the idiomatic patterns, build custom errors,
wrap and inspect them with `errors.Is`/`errors.As`, and learn when `panic`/
`recover` are (rarely) appropriate. This is what separates hobby code from
professional Go.

> **Recap from Session 10:** interfaces describe behavior. In fact, `error` *is*
> just an interface — which is why error handling is so flexible.

---

## 1. `error` is just an interface (10 min)

The built-in `error` type is a one-method interface:

```go
type error interface {
    Error() string
}
```

Anything with an `Error() string` method is an error. A function that can fail
returns an `error` as its **last** return value. `nil` means "no error" (success):

```go
func doThing() (Result, error) { ... }

result, err := doThing()
if err != nil {
    // handle the failure
    return
}
// use result — safe, because err was nil
```

> 🔑 **The golden pattern, again:** `if err != nil { ... }`. You met it in
> Session 05; now we go deep. Check errors *immediately*, handle them, and only
> use the result once you know `err == nil`.

### Creating simple errors

```go
import "errors"

errors.New("something went wrong")          // a basic error

import "fmt"
fmt.Errorf("user %d not found", id)          // a formatted error
```

`fmt.Errorf` is what you'll use most — it builds an error message with values
embedded.

Run [`examples/session11/basics/basics.go`](../examples/session11/basics/basics.go).

---

## 2. Handling errors well (10 min)

A few idioms that mark professional Go:

**Return early.** Handle the error and `return`, keeping the happy path
un-indented:

```go
f, err := os.Open("file.txt")
if err != nil {
    return err          // bail out now
}
defer f.Close()
// ... continue with f, knowing it's valid
```

**Add context as the error travels up.** Wrap with `%w` so callers learn *where*
it failed without losing the original:

```go
func loadConfig() error {
    data, err := os.ReadFile("config.json")
    if err != nil {
        return fmt.Errorf("loadConfig: %w", err)   // %w wraps the original
    }
    ...
}
```

Now the message reads like a trail: `loadConfig: open config.json: no such file`.

**Don't ignore errors.** `_ = doThing()` throws away a failure. Only ignore an
error when you genuinely don't care (and say so with a comment).

---

## 3. Custom error types & sentinel errors (20 min)

### Sentinel errors — predefined, comparable error values

When callers need to check for a *specific* error, define it once as a package
variable (convention: name starts with `Err`):

```go
var ErrNotFound = errors.New("not found")

func findUser(id int) (*User, error) {
    if id == 0 {
        return nil, ErrNotFound
    }
    ...
}

// Caller checks for that exact error with errors.Is:
user, err := findUser(0)
if errors.Is(err, ErrNotFound) {
    fmt.Println("no such user")
}
```

> 🔑 **Use `errors.Is`, not `==`.** Because errors get *wrapped* with `%w`,
> `errors.Is` walks the whole wrap chain to find a match. Plain `==` would miss a
> wrapped error.

### Custom error structs — errors that carry data

When an error needs to carry extra information, make a struct with an `Error()`
method:

```go
type ValidationError struct {
    Field string
    Msg   string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on %s: %s", e.Field, e.Msg)
}

func validate(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Msg: "must be non-negative"}
    }
    return nil
}
```

To pull the typed error back out (to read its fields), use `errors.As`:

```go
err := validate(-5)
var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Println("bad field:", ve.Field)   // "age"
}
```

> 🔑 **`errors.Is` vs `errors.As`:**
> - `errors.Is(err, target)` — "is this (anywhere in the chain) *this specific
>   error value*?" Use for sentinels.
> - `errors.As(err, &target)` — "is there an error of *this type* in the chain?
>   If so, extract it." Use for custom error structs whose fields you need.

Run [`examples/session11/custom/custom.go`](../examples/session11/custom/custom.go).

---

## 4. `panic` and `recover` — the emergency exit (15 min)

A **panic** is for *unrecoverable*, programmer-level problems — not for normal
errors. When a panic fires, the program unwinds the stack and crashes (unless
recovered).

```go
panic("this should never happen")
```

Things that panic automatically: out-of-bounds slice index, nil map write, nil
pointer dereference, divide by zero (integer).

> ⚠️ **Do NOT use panic for ordinary errors** like "file not found" or "invalid
> input." Return an `error` instead. Panic is for "the world is broken"
> situations or truly impossible states.

### `recover` — catching a panic

`recover` (only useful inside a `defer`ed function) stops a panic from crashing
the program. The main legitimate use is at a boundary you don't want to take
down the whole process — e.g. a web server recovering from a panic in one
request handler so the *server* keeps running:

```go
func safeRun(task func()) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("recovered from panic:", r)
        }
    }()
    task()   // even if this panics, we recover and continue
}
```

You'll use exactly this pattern as **recovery middleware** in the final project.

Run [`examples/session11/panic/panic.go`](../examples/session11/panic/panic.go).

> 💡 **Rule of thumb:** 99% of the time, return an `error`. Reserve `panic` for
> impossible states, and `recover` for protecting a boundary (like a server).

---

## 🎯 Exercises (do these before Session 12!)

Create `examples/session11/practice/practice.go`:

1. **Divide with error:** Write `func divide(a, b float64) (float64, error)` that
   returns an error when `b == 0`. Call it both ways and handle with `if err != nil`.
2. **Wrap with context:** Write a `func parseAge(s string) (int, error)` that uses
   `strconv.Atoi` and, on failure, returns `fmt.Errorf("parseAge: %w", err)`.
   Print the wrapped message.
3. **Sentinel error:** Define `var ErrEmptyName = errors.New("name is empty")`.
   Write `func greet(name string) (string, error)` returning it for `""`. In the
   caller, detect it with `errors.Is`.
4. **Custom error type:** Build a `RangeError struct { Value, Min, Max int }` with
   an `Error()` method. Write a `func check(n int) error` returning it when out of
   range, and use `errors.As` to read `.Min`/`.Max` in the caller.
5. **Recover:** Write `func safeDivideInts(a, b int) (result int, err error)` that
   does `a / b` but uses `defer`+`recover` to turn the divide-by-zero panic into a
   returned error instead of a crash.

---

## ✅ Session 11 Checklist

- [ ] I know `error` is a one-method interface and `nil` means success
- [ ] I use `errors.New` and `fmt.Errorf` to create errors
- [ ] I check errors immediately and return early
- [ ] I wrap errors with `%w` to add context
- [ ] I define sentinel errors and detect them with `errors.Is`
- [ ] I build custom error structs and extract them with `errors.As`
- [ ] I know panic is for unrecoverable situations, not normal errors
- [ ] I can recover from a panic inside a deferred function
- [ ] I completed all 5 exercises

**Previous:** [← Session 10](session-10.md) · **Next:** [Session 12 — Concurrency I →](session-12.md)
