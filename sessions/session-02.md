> 🌐 **Language / زبان:** English (this file) · [فارسی](session-02.fa.md)

# Session 02 — Variables & Types 🧮

**Goal (1 hour):** Learn how Go stores data. You'll master the two ways to
declare variables, understand Go's basic types, learn the all-important concept
of **zero values**, work with constants, and convert between types.

> **Recap from Session 01:** every runnable program has `package main` + `func main()`,
> and we print with `fmt.Println`/`Printf`. We'll use those constantly now.

---

## 1. What is a variable? (5 min)

A **variable** is a named box that holds a value. In Go, every variable has a
**type** (string, integer, etc.) that never changes after it's declared. Go is
**statically typed** — the type is fixed at compile time, which catches whole
classes of bugs before your program ever runs.

---

## 2. Declaring variables — the two ways (15 min)

### Way 1: `var` (explicit)

```go
var name string = "Sobhan"
var age int = 25
```

You can let Go **infer** the type from the value (no need to write `string`):

```go
var name = "Sobhan"   // Go sees a string, so name is a string
var age = 25          // Go sees a whole number, so age is an int
```

You can also declare *without* a value — it gets the **zero value** (next section):

```go
var count int       // count is 0
var message string  // message is "" (empty string)
```

### Way 2: `:=` (short declaration) — the one you'll use most

Inside a function, `:=` declares **and** assigns in one step, with type inference:

```go
name := "Sobhan"   // same as: var name string = "Sobhan"
age := 25          // same as: var age int = 25
```

> 🔑 **Two key rules for `:=`**
> 1. It only works **inside a function** (not at package/global level — use `var` there).
> 2. The variable must be **new**. To just reassign an existing variable, use `=`:
> ```go
> age := 25   // declare
> age = 26    // reassign (no colon)
> ```

### Declaring several at once

```go
var (
    firstName = "Sobhan"
    lastName  = "Azimzadeh"
    age       = 25
)

// or with :=
x, y := 10, 20
```

Run [`examples/session02/variables/variables.go`](../examples/session02/variables/variables.go):

```bash
go run examples/session02/variables/variables.go
```

---

## 3. The basic types (15 min)

Go's built-in types you'll use daily:

| Category | Types | Example |
|----------|-------|---------|
| **String** | `string` | `"hello"` |
| **Integer** | `int`, `int8/16/32/64`, `uint`, `uint8/16/32/64` | `42`, `-7` |
| **Float** | `float32`, `float64` | `3.14`, `-0.5` |
| **Boolean** | `bool` | `true`, `false` |
| **Byte/Rune** | `byte` (=uint8), `rune` (=int32) | covered in Session 07 |

**Practical advice for beginners — just use these defaults:**
- Whole numbers → `int`
- Decimal numbers → `float64`
- Text → `string`
- True/false → `bool`

The sized variants (`int8`, `uint32`, …) matter when you care about memory or
binary formats. You'll know when you need them. `int` is 64-bit on modern machines.

```go
var temperature float64 = 36.6
var isHealthy bool = true
var city string = "Tehran"
var population int = 9_500_000   // underscores allowed for readability
```

See [`examples/session02/types/types.go`](../examples/session02/types/types.go).

---

## 4. Zero values — Go has no "undefined" (10 min)

This is one of Go's best features. **Every variable is always initialized.**
If you don't give a value, Go gives it the **zero value** for its type:

| Type | Zero value |
|------|-----------|
| `int`, `float64` | `0` |
| `string` | `""` (empty string) |
| `bool` | `false` |
| pointers, slices, maps, etc. | `nil` (later sessions) |

```go
var count int       // 0, not garbage, not undefined
var name string     // ""
var active bool     // false
```

> 🔑 **Why it matters:** In many languages an uninitialized variable holds
> garbage or `undefined`, causing crashes. In Go, a freshly declared variable is
> always safe to use. No null-pointer surprises on basic types.

---

## 5. Constants (5 min)

A **constant** is a value that can never change. Declare with `const`:

```go
const Pi = 3.14159
const AppName = "TaskFlow"
const MaxUsers = 100
```

```go
const (
    StatusActive   = "active"
    StatusInactive = "inactive"
)
```

Use constants for values that are fixed forever (config names, math constants,
fixed limits). Trying to reassign a constant is a compile error — which is
exactly what you want.

---

## 6. Type conversion — Go never converts for you (10 min)

Go is strict: it will **not** automatically mix types. This is a compile error:

```go
var x int = 5
var y float64 = 2.5
// z := x + y   // ❌ ERROR: mismatched types int and float64
```

You must **explicitly convert** using `T(value)`:

```go
var x int = 5
var y float64 = 2.5
z := float64(x) + y   // ✅ convert x to float64 first → 7.5
```

```go
a := 10
b := 3
fmt.Println(a / b)              // 3   (integer division — truncates!)
fmt.Println(float64(a) / float64(b)) // 3.3333... (float division)
```

> ⚠️ **Classic beginner trap:** `10 / 3` is `3`, not `3.33`, because both are
> integers and integer division throws away the remainder. Convert to `float64`
> when you want decimals.

Converting numbers ↔ strings needs the `strconv` package (Session 14), **not** a
plain conversion — `string(65)` gives `"A"` (a character), not `"65"`!

See [`examples/session02/conversion/conversion.go`](../examples/session02/conversion/conversion.go).

---

## 🎯 Exercises (do these before Session 03!)

Create `examples/session02/practice/practice.go` (`package main`, `func main()`):

1. **Profile card:** Declare your `name` (string), `age` (int), `height`
   (float64), and `student` (bool) using **`:=`**. Print them with `Printf`,
   one per line, using the right verb for each.
2. **Zero values:** Declare `var a int`, `var b string`, `var c bool` with **no
   values**, and print them. Predict the output first, then verify.
3. **The division trap:** Print `7 / 2` and then `float64(7) / float64(2)`.
   Explain to yourself (in a comment) why they differ.
4. **Convert:** You have `var meters float64 = 1.81`. Convert it to an `int` and
   print both. What happened to the decimal part? (This teaches truncation.)
5. **Constants:** Make a `const TaxRate = 0.09`. Given `price := 100.0`, compute
   and print the final price including tax.

---

## ✅ Session 02 Checklist

- [ ] I can declare variables with both `var` and `:=`
- [ ] I know `:=` is function-only and must declare something new
- [ ] I can name the zero values for int, string, and bool
- [ ] I know Go never auto-converts types — I must convert explicitly
- [ ] I understand why `7/2 == 3` for integers
- [ ] I can declare constants with `const`
- [ ] I completed all 5 exercises

**Previous:** [← Session 01](session-01.md) · **Next:** [Session 03 — Control Flow →](session-03.md)
