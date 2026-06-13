> 🌐 **Language / زبان:** English (this file) · [فارسی](session-10.fa.md)

# Session 10 — Interfaces 🔌

**Goal (1 hour):** Master Go's signature feature. An **interface** describes
*what something can do* (its behavior) without saying *what it is*. Interfaces
give you flexible, decoupled code — and Go's twist (they're satisfied
**implicitly**) is what makes them so pleasant. This is the session that turns a
beginner into a confident Go developer.

> **Recap from Session 09:** structs + methods + pointers let you model data and
> attach behavior. Interfaces let different types be used *interchangeably* based
> on shared behavior.

---

## 1. What is an interface? (10 min)

An **interface** is a set of method signatures. Any type that has those methods
*automatically* satisfies the interface — **no `implements` keyword needed.**

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

This says: "a `Shape` is anything that has an `Area() float64` method **and** a
`Perimeter() float64` method." That's it. Any type with both methods *is* a
`Shape`, automatically.

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

// Circle now satisfies Shape — we never declared it does. It just does.
var s Shape = Circle{Radius: 5}
fmt.Println(s.Area())
```

> 🔑 **Implicit satisfaction is Go's superpower.** In Java/C# you must write
> `class Circle implements Shape`. In Go, if the methods match, it just works.
> This means you can make *existing* types (even from other packages) satisfy
> *your* interfaces.

---

## 2. Polymorphism: one function, many types (20 min)

Because any matching type is a `Shape`, you can write functions that work on the
*interface* and accept any concrete type that satisfies it:

```go
func describe(s Shape) {
    fmt.Printf("area=%.2f perimeter=%.2f\n", s.Area(), s.Perimeter())
}

describe(Circle{Radius: 5})
describe(Rectangle{Width: 3, Height: 4})
```

The same `describe` works for circles, rectangles, triangles — anything that
"is a Shape." You can even hold mixed types in one slice:

```go
shapes := []Shape{
    Circle{Radius: 2},
    Rectangle{Width: 3, Height: 4},
}
for _, s := range shapes {
    describe(s)
}
```

This is **polymorphism** — writing code against behavior, not concrete types.
It's the key to flexible, testable Go: your functions depend on small interfaces,
and you can swap in any implementation (real one, or a fake for testing).

Run [`examples/session10/basics/basics.go`](../examples/session10/basics/basics.go) and
[`examples/session10/polymorphism/polymorphism.go`](../examples/session10/polymorphism/polymorphism.go).

---

## 3. The `Stringer` interface — make your types print nicely (10 min)

The standard library defines tiny interfaces you'll satisfy all the time. The
most common is `fmt.Stringer`:

```go
type Stringer interface {
    String() string
}
```

If your type has a `String() string` method, `fmt` automatically uses it when
printing — exactly what you saw with the `Status` enum in Session 03:

```go
type Color struct{ R, G, B int }

func (c Color) String() string {
    return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

fmt.Println(Color{255, 0, 0})   // rgb(255, 0, 0)  — uses String()!
```

> 💡 **Design principle:** Go favors **small interfaces**. The most famous,
> `io.Writer` and `io.Reader`, each have *one* method. Small interfaces are easy
> to satisfy and easy to mock. "The bigger the interface, the weaker the
> abstraction" — a Go proverb.

---

## 4. The empty interface & type assertions (15 min)

The empty interface `interface{}` (or its alias **`any`**, since Go 1.18) has
*no* methods — so **every** type satisfies it. It means "any value at all":

```go
var x any        // can hold anything
x = 42
x = "hello"
x = []int{1, 2, 3}
```

You'll see `any` in functions that accept arbitrary values (like `fmt.Println`).
But once a value is stored as `any`, you've lost its concrete type. To get it
back, use a **type assertion**:

```go
var x any = "hello"

s := x.(string)       // assert x is a string -> "hello"
fmt.Println(s)

// Safe form with comma-ok (no panic if wrong):
if s, ok := x.(string); ok {
    fmt.Println("it's a string:", s)
}
```

### The type switch — handle many possible types

```go
func describe(i any) {
    switch v := i.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string of length", len(v))
    case bool:
        fmt.Println("bool:", v)
    default:
        fmt.Printf("unknown type %T\n", v)
    }
}
```

> ⚠️ **Use `any` sparingly.** Reaching for `any` often means you're throwing away
> the type safety Go gives you. Prefer real types and small interfaces; use `any`
> only when you genuinely must accept "anything" (like generic printers or JSON).

Run [`examples/session10/empty/empty.go`](../examples/session10/empty/empty.go).

---

## 🎯 Exercises (do these before Session 11!)

Create `examples/session10/practice/practice.go`:

1. **Shape interface:** Define `Shape` with `Area() float64`. Implement it for
   `Circle` and `Square`. Write `func totalArea(shapes []Shape) float64` that
   sums the areas of a mixed slice. Test with a few shapes.
2. **Stringer:** Define a `Money struct { Amount float64; Currency string }` with
   a `String()` method that prints like `"$19.99"`. Print a `Money` value and
   confirm `fmt` uses your method.
3. **Speaker:** Define `type Speaker interface { Speak() string }`. Implement
   `Dog` and `Cat`. Write a function that takes `[]Speaker` and prints what each
   one says.
4. **Type switch:** Write `func printType(values ...any)` that uses a type switch
   to report each argument's type and value. Call it with an int, a string, a
   float, and a bool in one call.
5. **Mock for testing (preview):** Define `type Notifier interface { Send(msg
   string) error }`. Make a real `EmailNotifier` and a `MockNotifier` that just
   records messages. Write a function that takes a `Notifier` — see how you could
   test it without sending real emails. *(This is exactly how interfaces enable
   testing in the final project!)*

---

## ✅ Session 10 Checklist

- [ ] I can define an interface as a set of method signatures
- [ ] I understand satisfaction is implicit — no `implements` keyword
- [ ] I can write a function that accepts an interface and works on many types
- [ ] I can store different concrete types in a `[]Interface` slice
- [ ] I can implement `String()` to control how a type prints
- [ ] I know `any` = `interface{}` and how to recover the type with assertions
- [ ] I can write a type switch
- [ ] I know to prefer small interfaces and use `any` sparingly
- [ ] I completed all 5 exercises

**Previous:** [← Session 09](session-09.md) · **Next:** [Session 11 — Errors →](session-11.md)
