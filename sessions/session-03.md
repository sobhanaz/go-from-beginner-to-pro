> 🌐 **Language / زبان:** English (this file) · [فارسی](session-03.fa.md)

# Session 03 — Control Flow 🔀

**Goal (1 hour):** Teach your programs to make decisions. You'll master `if`/`else`,
Go's powerful `switch`, the operators that drive them, and the `iota` trick for
clean constants.

> **Recap from Session 02:** you can declare variables (`:=`, `var`), you know the
> four basic types and their zero values, and you know Go never auto-converts types.

---

## 1. Operators — the building blocks of decisions (10 min)

Decisions are built from expressions that evaluate to `true` or `false`.

**Comparison operators** (result is always a `bool`):

| Operator | Meaning | Example |
|----------|---------|---------|
| `==` | equal to | `age == 18` |
| `!=` | not equal to | `name != ""` |
| `<` `>` | less / greater than | `score > 50` |
| `<=` `>=` | less/greater or equal | `age >= 18` |

**Logical operators** (combine booleans):

| Operator | Meaning | True when… |
|----------|---------|-----------|
| `&&` | AND | **both** sides are true |
| `\|\|` | OR | **at least one** side is true |
| `!` | NOT | flips true↔false |

```go
age := 20
hasTicket := true

canEnter := age >= 18 && hasTicket   // true
fmt.Println(canEnter)
fmt.Println(!hasTicket)              // false
```

**Arithmetic** you already met: `+ - * /` and `%` (modulo = remainder).
`%` is incredibly useful: `n % 2 == 0` tests if `n` is even.

---

## 2. `if` / `else if` / `else` (15 min)

The basic shape — **note: no parentheses around the condition**, but braces are
**required** even for one line:

```go
score := 75

if score >= 90 {
    fmt.Println("A")
} else if score >= 70 {
    fmt.Println("B")
} else {
    fmt.Println("Needs improvement")
}
```

> 🔑 Go style: the opening `{` is always on the same line as the `if`. This is not
> optional — `gofmt` enforces it and the compiler expects it.

### The `if` with a short statement (very idiomatic Go!)

You can run a little statement *before* the condition, separated by `;`. The
variable it creates is **only visible inside the `if`/`else`** — this keeps scope tight:

```go
if age := getAge(); age >= 18 {
    fmt.Println("Adult, age is", age)
} else {
    fmt.Println("Minor, age is", age)
}
// age does NOT exist out here
```

You'll see this pattern *everywhere* in real Go, especially with errors:

```go
if err := doSomething(); err != nil {
    // handle the error
}
```

Run [`examples/session03/ifelse/ifelse.go`](../examples/session03/ifelse/ifelse.go).

---

## 3. `switch` — cleaner than long if/else chains (20 min)

When you compare one value against many options, `switch` is clearer:

```go
day := "Sat"

switch day {
case "Sat", "Sun":          // multiple values per case!
    fmt.Println("Weekend 🎉")
case "Mon":
    fmt.Println("Monday...")
default:
    fmt.Println("A weekday")
}
```

**Two things that surprise people coming from other languages:**

1. **No `break` needed.** Go automatically stops after a matching case — there's
   no accidental "fall-through" into the next case.
2. **One case can list multiple values** separated by commas (`case "Sat", "Sun":`).

### `switch` with no condition (replaces if/else chains)

Leave out the value and each `case` becomes a boolean test. This is the cleanest
way to write a range/threshold check:

```go
switch {
case score >= 90:
    fmt.Println("A")
case score >= 70:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

### `switch` with a short statement

Just like `if`, you can initialize a variable first:

```go
switch hour := time.Now().Hour(); {
case hour < 12:
    fmt.Println("Good morning")
case hour < 18:
    fmt.Println("Good afternoon")
default:
    fmt.Println("Good evening")
}
```

> 💡 `fallthrough` exists if you *really* want to continue into the next case, but
> it's rare. Mention it in interviews; you'll seldom use it.

Run [`examples/session03/switch/switch.go`](../examples/session03/switch/switch.go).

---

## 4. `iota` — auto-numbering constants (10 min)

When you need a set of related constants (like statuses or levels), typing
`0, 1, 2, 3` by hand is error-prone. `iota` auto-increments inside a `const` block:

```go
const (
    Sunday = iota   // 0
    Monday          // 1
    Tuesday         // 2
    Wednesday       // 3
)
```

`iota` starts at `0` on the first line of the `const` block and increases by 1
each line. This is the idiomatic way to make **enumerations** (enums) in Go.

A common pro pattern uses a named type for safety:

```go
type Status int

const (
    StatusPending Status = iota   // 0
    StatusActive                  // 1
    StatusClosed                  // 2
)
```

Now a `Status` value can only be one of those three — your code is self-documenting.
You'll use this in the final project for things like task status.

Run [`examples/session03/iota/iota.go`](../examples/session03/iota/iota.go).

---

## 🎯 Exercises (do these before Session 04!)

Create `examples/session03/practice/practice.go`:

1. **Grade calculator:** Given `score := 83`, use `if/else if/else` to print the
   letter grade (90+ = A, 80+ = B, 70+ = C, else F).
2. **Even or odd:** Given `n := 17`, use the `%` operator and an `if` to print
   whether it's even or odd.
3. **Rewrite with switch:** Redo exercise 1 using a **condition-less `switch`**.
   Notice how much cleaner it reads.
4. **FizzBuzz (the classic interview question!):** Print numbers 1 to 20, but for
   multiples of 3 print "Fizz", multiples of 5 print "Buzz", and multiples of
   both print "FizzBuzz". *(You'll need a loop — peek ahead, or just do 1–20 by
   hand with if/switch. We cover loops properly next session.)*
5. **iota enum:** Make a `type Level int` with `Low`, `Medium`, `High` via `iota`,
   then print each one's numeric value.

---

## ✅ Session 03 Checklist

- [ ] I can use `==`, `!=`, `<`, `&&`, `||`, `!` correctly
- [ ] I write `if` with no parentheses but always with braces
- [ ] I can use the `if x := ...; cond {` short-statement form
- [ ] I know Go `switch` needs no `break` and allows multiple values per case
- [ ] I can write a condition-less `switch` to replace an if/else chain
- [ ] I can create an enum with `iota`
- [ ] I completed all 5 exercises

**Previous:** [← Session 02](session-02.md) · **Next:** [Session 04 — Loops & Functions →](session-04.md)
