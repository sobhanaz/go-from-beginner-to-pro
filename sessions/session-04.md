> 🌐 **Language / زبان:** English (this file) · [فارسی](session-04.fa.md)

# Session 04 — Loops & Functions 🔁

**Goal (1 hour):** Learn to repeat work with the `for` loop (Go's *only* loop —
it does everything) and to organize code into reusable **functions**. These two
tools are the backbone of every program you'll write.

> **Recap from Session 03:** you can branch with `if`/`switch` and combine
> conditions with `&&`, `||`, `!`. Now we add repetition and reusability.

---

## 1. The `for` loop — Go's only loop (20 min)

Many languages have `for`, `while`, `do-while`. **Go has just `for`** — and it
covers every case. Learn its shapes:

### Shape 1: the classic three-part loop

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)   // prints 0 1 2 3 4
}
```

Three parts separated by `;`:
1. **init**: `i := 0` — runs once at the start.
2. **condition**: `i < 5` — checked before each iteration; loop runs while true.
3. **post**: `i++` — runs after each iteration. (`i++` means `i = i + 1`.)

### Shape 2: the "while" loop (just a condition)

Drop the init and post — now it behaves like a `while`:

```go
n := 1
for n < 100 {
    n = n * 2
}
fmt.Println(n)   // 128
```

### Shape 3: the infinite loop

```go
for {
    // runs forever until you break out
    break   // exits the loop
}
```

You'll use this for servers and event loops. Use `break` to exit and `continue`
to skip to the next iteration:

```go
for i := 0; i < 10; i++ {
    if i == 3 {
        continue   // skip 3
    }
    if i == 6 {
        break      // stop entirely at 6
    }
    fmt.Println(i) // 0 1 2 4 5
}
```

### Shape 4: `for range` — looping over collections

The most common form in real code. It walks through items of a slice, string,
map, etc. (we'll use it heavily once we hit slices in Session 06):

```go
names := []string{"Ann", "Bob", "Cy"}
for index, name := range names {
    fmt.Println(index, name)   // 0 Ann / 1 Bob / 2 Cy
}
```

If you don't need the index, use the blank identifier `_`:

```go
for _, name := range names {
    fmt.Println(name)
}
```

> 🔑 **`_` (underscore)** is the "I don't want this value" placeholder. Go forces
> you to handle every value, but `_` lets you politely discard one.

Run [`examples/session04/loops/loops.go`](../examples/session04/loops/loops.go).

---

## 2. Functions — reusable blocks of logic (20 min)

A **function** packages code so you can call it by name, give it inputs
(**parameters**), and get back an output (**return value**).

### Anatomy

```go
func add(a int, b int) int {
    return a + b
}
```

| Part | Meaning |
|------|---------|
| `func` | keyword that starts a function |
| `add` | the function's name |
| `(a int, b int)` | parameters with their types |
| `int` (after the `)`) | the **return type** |
| `return a + b` | sends a value back to the caller |

Call it like this:

```go
sum := add(3, 4)   // sum == 7
```

### Shorthand for same-type parameters

If consecutive parameters share a type, write the type once:

```go
func add(a, b int) int {   // same as (a int, b int)
    return a + b
}
```

### Functions with no return

If a function returns nothing, omit the return type:

```go
func greet(name string) {
    fmt.Println("Hello,", name)
}
```

### Where can you call a function?

Anywhere in the same package — **order doesn't matter**. You can define `main`
at the top and helper functions below it; Go finds them. (You saw this in
Session 03 with `getAge`.)

Run [`examples/session04/functions/functions.go`](../examples/session04/functions/functions.go).

---

## 3. Putting it together: FizzBuzz done right (10 min)

Now you have everything to solve the classic interview question properly with a
loop. The full, idiomatic solution is in
[`examples/session04/fizzbuzz/fizzbuzz.go`](../examples/session04/fizzbuzz/fizzbuzz.go):

```go
func fizzbuzz(n int) {
    for i := 1; i <= n; i++ {
        switch {
        case i%15 == 0:
            fmt.Println("FizzBuzz")
        case i%3 == 0:
            fmt.Println("Fizz")
        case i%5 == 0:
            fmt.Println("Buzz")
        default:
            fmt.Println(i)
        }
    }
}
```

Notice how `for`, `switch`, `%`, and a function come together. **Why check `%15`
first?** Because a number divisible by both 3 and 5 is divisible by 15 — and the
condition-less `switch` stops at the first match, so the most specific case must
come first.

```bash
go run examples/session04/fizzbuzz/fizzbuzz.go
```

---

## 🎯 Exercises (do these before Session 05!)

Create `examples/session04/practice/practice.go`:

1. **Sum 1..100:** Use a `for` loop to add up every number from 1 to 100 and
   print the total. (The famous answer is 5050.)
2. **Countdown:** Loop from 10 down to 1 (`i--`), printing each, then print
   "Liftoff!".
3. **Function `isPrime`:** Write `func isPrime(n int) bool` that returns whether
   `n` is prime. Test it on 7, 10, and 13.
4. **Function `factorial`:** Write `func factorial(n int) int` using a loop that
   returns `n!` (e.g. `factorial(5)` → 120).
5. **Combine:** Write `func sumEven(limit int) int` that returns the sum of all
   even numbers from 1 to `limit`, using `continue` to skip odd numbers.

---

## ✅ Session 04 Checklist

- [ ] I know Go has only one loop keyword: `for`
- [ ] I can write the 3-part, "while", and infinite forms of `for`
- [ ] I can use `break` and `continue`
- [ ] I understand `for range` and the `_` blank identifier
- [ ] I can define a function with parameters and a return type
- [ ] I can call functions and use their return values
- [ ] I solved FizzBuzz with a loop
- [ ] I completed all 5 exercises

**Previous:** [← Session 03](session-03.md) · **Next:** [Session 05 — Functions Deep Dive →](session-05.md)
