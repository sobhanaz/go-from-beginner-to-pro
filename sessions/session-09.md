> 🌐 **Language / زبان:** English (this file) · [فارسی](session-09.fa.md)

# Session 09 — Pointers 📍

**Goal (1 hour):** Demystify pointers once and for all. You'll learn what a
pointer *is*, the two operators `&` and `*`, why functions need pointers to
modify their arguments, and when (and when not) to use them. Pointers sound
scary; in Go they're simple and safe.

> **Recap from Session 08:** you used pointer *receivers* (`func (r *T)`) to let
> methods modify a struct. This session explains the `*` and `&` behind that.

---

## 1. What is a pointer? (10 min)

Every value your program stores lives at some **address** in memory. A
**pointer** is just a value that holds that address — it "points to" where
another value lives.

An analogy: a variable is a house, its value is what's inside, and a pointer is
the *street address* written on a piece of paper. With the address you can go
find the house and change what's inside.

```go
x := 10        // a normal int
p := &x        // p is a pointer to x (holds x's address)

fmt.Println(x)   // 10   — the value
fmt.Println(p)   // 0xc0000140a8 — an address (yours will differ)
fmt.Println(*p)  // 10   — follow the pointer to get the value
```

Two operators are all you need:

| Operator | Name | What it does |
|----------|------|--------------|
| `&x` | "address of" | gives you a **pointer to** `x` |
| `*p` | "dereference" | follows the pointer to get/set the **value** it points to |

The **type** of a pointer to an `int` is written `*int`. A pointer to a `User`
is `*User`.

---

## 2. Reading and writing through a pointer (10 min)

Dereferencing with `*p` works both ways — you can read *and* assign:

```go
x := 10
p := &x

*p = 20          // change the value AT the address p points to
fmt.Println(x)   // 20  ← x changed, because p points to x!
```

This is the whole point of pointers: two names (`x` and `*p`) refer to the
**same** piece of memory, so changing one changes the other.

### The zero value of a pointer is `nil`

A pointer that points to nothing is `nil`. Dereferencing a `nil` pointer
**panics** (crashes), so check before using one that might be nil:

```go
var p *int        // nil — points to nothing
fmt.Println(p)    // <nil>
// fmt.Println(*p) // 💥 PANIC: nil pointer dereference
```

Run [`examples/session09/basics/basics.go`](../examples/session09/basics/basics.go).

---

## 3. Why pointers matter: modifying function arguments (20 min)

This is the #1 reason beginners need pointers. **Go passes arguments by value —
it copies them.** So a function that takes a plain `int` gets a *copy* and can't
change the caller's variable:

```go
func tryToDouble(n int) {
    n = n * 2          // changes the COPY only
}

x := 5
tryToDouble(x)
fmt.Println(x)   // still 5 — the original was never touched
```

To actually modify the caller's variable, pass a **pointer** to it:

```go
func double(n *int) {
    *n = *n * 2        // follow the pointer, change the real value
}

x := 5
double(&x)             // pass the ADDRESS of x
fmt.Println(x)   // 10 — it worked!
```

> 🔑 **The pattern:** function parameter is `*T`, you call it with `&variable`,
> and inside you use `*param` to read/write the real value.

This is exactly why pointer *receivers* on methods (Session 08) can modify a
struct: the receiver is secretly a pointer to the original.

Run [`examples/session09/funcparams/funcparams.go`](../examples/session09/funcparams/funcparams.go).

---

## 4. Pointers to structs — the common case (15 min)

In real Go you mostly use pointers with **structs**, for two reasons: to modify
them in functions, and to avoid copying big structs around.

```go
type User struct {
    Name string
    Age  int
}

func birthday(u *User) {
    u.Age++          // Go lets you write u.Age, not (*u).Age — nice shorthand!
}

user := User{Name: "Sobhan", Age: 25}
birthday(&user)
fmt.Println(user.Age)   // 26
```

> 💡 **Convenience:** for a struct pointer `u`, Go automatically dereferences for
> field access — `u.Age` means `(*u).Age`. You almost never write `(*u)` yourself.

### `new` and pointer creation

You can make a pointer to a fresh zero-valued struct two ways:

```go
u1 := &User{Name: "Ali"}   // most common: address of a struct literal
u2 := new(User)            // new(T) returns a *T pointing to a zero value
u2.Name = "Sara"
```

`&User{...}` is by far the most common; `new` shows up occasionally.

Run [`examples/session09/structs/structs.go`](../examples/session09/structs/structs.go).

---

## 5. When should you use pointers? (5 min)

Don't sprinkle `*` everywhere. Use a pointer when:

- ✅ A function/method needs to **modify** its argument or receiver.
- ✅ The struct is **large** and copying it would be wasteful.
- ✅ You need to represent "**no value**" with `nil` (e.g. an optional field).

Prefer plain values when:

- ✅ The data is small (an `int`, a small struct) and read-only.
- ✅ You want safety and simplicity — value copies can't be mutated by surprise.

> 🔑 **Good news:** Go has a **garbage collector**, so you never manually free
> memory. No `malloc`/`free`, no leaks from forgetting to free. And there's no
> pointer *arithmetic* like in C — Go pointers are safe and simple.

---

## 🎯 Exercises (do these before Session 10!)

Create `examples/session09/practice/practice.go`:

1. **Swap:** Write `func swap(a, b *int)` that swaps the two values the pointers
   point to. Test it: declare `x, y := 1, 2`, call `swap(&x, &y)`, print both.
2. **Increment:** Write `func increment(n *int)` that adds 1 to the value. Call
   it in a loop 5 times on the same variable and print the result (should be 5
   if it started at 0).
3. **Modify a struct:** Define `type Account struct { Balance float64 }`. Write
   `func deposit(a *Account, amount float64)` that increases the balance. Prove
   the original account changed.
4. **Value vs pointer:** Write two functions — `func resetValue(a Account)` and
   `func resetPointer(a *Account)` — both setting `Balance = 0`. Show which one
   actually affects the caller's account, and explain why in a comment.
5. **Nil check:** Write `func describe(p *int)` that prints the value if `p` is
   not nil, otherwise prints `"no value"`. Call it with both `&someInt` and a
   plain `nil`.

---

## ✅ Session 09 Checklist

- [ ] I can explain that a pointer holds a memory address
- [ ] I know `&x` gives a pointer and `*p` dereferences it
- [ ] I understand Go copies arguments, so I need `*T` + `&x` to modify them
- [ ] I can write a function that modifies its argument via a pointer
- [ ] I know `u.Field` auto-dereferences for struct pointers
- [ ] I know a nil pointer dereference panics, and I check for nil
- [ ] I know when to prefer pointers vs values
- [ ] I completed all 5 exercises

**Previous:** [← Session 08](session-08.md) · **Next:** [Session 10 — Interfaces →](session-10.md)

---

🎉 **Milestone:** Part 2 (Data Structures) complete! You can now model any data:
slices, maps, structs, and pointers. Next comes Part 3 — the ideas that make Go
*Go*: interfaces, errors, and concurrency.
