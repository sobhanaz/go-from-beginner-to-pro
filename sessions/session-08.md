> 🌐 **Language / زبان:** English (this file) · [فارسی](session-08.fa.md)

# Session 08 — Structs & Methods 🏗️

**Goal (1 hour):** Define your own types. A **struct** groups related fields into
one custom type (like a `User` or a `Task`), and **methods** attach behavior to
those types. This is how you model the real world in Go — and it's the backbone
of the final project.

> **Recap from Session 07:** maps and strings round out the built-in data types.
> Now you build your *own* types to represent things in your program's domain.

---

## 1. Structs — grouping related data (15 min)

A **struct** is a typed collection of named **fields**. Define one with `type`:

```go
type User struct {
    ID    int
    Name  string
    Email string
    Active bool
}
```

### Creating struct values

```go
// 1. Named fields (best — clear and order-independent).
u1 := User{
    ID:    1,
    Name:  "Sobhan",
    Email: "sobhan@example.com",
    Active: true,
}

// 2. Positional (must match field order — avoid for big structs).
u2 := User{2, "Ali", "ali@example.com", false}

// 3. Zero value — every field gets its type's zero value.
var u3 User   // {0 "" "" false}
```

### Accessing and changing fields

```go
fmt.Println(u1.Name)   // Sobhan
u1.Name = "Sobhan A."  // fields are mutable
u1.Active = false
```

> 💡 **Field naming = visibility.** A field starting with an **uppercase** letter
> (`Name`) is **exported** (visible from other packages). Lowercase (`name`) is
> **unexported** (private to the package). This capitalization rule applies to
> everything in Go — types, functions, fields. Remember it!

Run [`examples/session08/structs/structs.go`](../examples/session08/structs/structs.go).

---

## 2. Methods — behavior attached to a type (20 min)

A **method** is a function with a special **receiver** that ties it to a type.
The receiver goes in parentheses *before* the method name:

```go
type Rectangle struct {
    Width, Height float64
}

// (r Rectangle) is the receiver. Call as rect.Area().
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

rect := Rectangle{Width: 3, Height: 4}
fmt.Println(rect.Area())   // 12
```

### Value receiver vs. pointer receiver — the crucial distinction

```go
// VALUE receiver: gets a COPY. Changes don't affect the original.
func (r Rectangle) Scale(factor float64) {
    r.Width *= factor   // modifies the copy only — useless here!
}

// POINTER receiver: gets the actual struct. Changes persist.
func (r *Rectangle) ScaleInPlace(factor float64) {
    r.Width *= factor   // modifies the real struct
    r.Height *= factor
}
```

> 🔑 **The rule of thumb:**
> - Use a **pointer receiver** `(r *T)` when the method needs to **modify** the
>   struct, or the struct is large (avoid copying).
> - Use a **value receiver** `(r T)` for small structs that you only read.
> - **Be consistent:** if any method needs a pointer receiver, use pointer
>   receivers for *all* methods on that type. (Pointers are Session 09 — for now,
>   just know `*T` lets you change the original.)

Go conveniently lets you call pointer methods on addressable values:
`rect.ScaleInPlace(2)` works even though `rect` isn't explicitly a pointer.

### Methods can be on any named type, not just structs

```go
type Celsius float64

func (c Celsius) ToFahrenheit() float64 {
    return float64(c)*9/5 + 32
}

temp := Celsius(100)
fmt.Println(temp.ToFahrenheit())   // 212
```

(You saw a sneak peek of this with the `String()` method on the `Status` enum
back in Session 03.)

Run [`examples/session08/methods/methods.go`](../examples/session08/methods/methods.go).

---

## 3. Constructors — the `New...` convention (10 min)

Go has no `class` and no built-in constructor. The idiom is a plain function
named `NewX` that returns a ready-to-use value:

```go
func NewUser(name, email string) User {
    return User{
        Name:   name,
        Email:  email,
        Active: true,   // sensible default
    }
}

u := NewUser("Sara", "sara@example.com")
```

This is just a convention, but it's everywhere in real Go. Use it to enforce
defaults and validation when creating your types.

---

## 4. Struct embedding — composition over inheritance (15 min)

Go has **no inheritance**. Instead it has **embedding**: put one struct inside
another *without a field name*, and its fields/methods get "promoted" to the outer
struct.

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " makes a sound"
}

type Dog struct {
    Animal       // embedded — no field name
    Breed string
}

d := Dog{
    Animal: Animal{Name: "Rex"},
    Breed:  "Labrador",
}

fmt.Println(d.Name)      // "Rex"  — promoted from Animal!
fmt.Println(d.Speak())   // "Rex makes a sound" — method promoted too
fmt.Println(d.Breed)     // "Labrador"
```

This is Go's "composition over inheritance" philosophy: you build bigger types
by *combining* smaller ones, not by extending a base class. It's cleaner and
avoids deep inheritance hierarchies.

Run [`examples/session08/embedding/embedding.go`](../examples/session08/embedding/embedding.go).

---

## 🎯 Exercises (do these before Session 09!)

Create `examples/session08/practice/practice.go`:

1. **Define a `Task`:** Make a struct `Task` with fields `ID int`, `Title string`,
   `Done bool`, `Priority int`. Create one with named fields and print it with
   `%+v` (shows field names).
2. **Method `Complete`:** Add a **pointer-receiver** method `func (t *Task)
   Complete()` that sets `Done = true`. Create a task, call it, and print to
   confirm the change stuck.
3. **Method `String`:** Add `func (t Task) String() string` returning a nice
   one-line summary (e.g. `"[✓] Buy milk (priority 2)"`). Print the task and
   watch Go use your `String()` automatically.
4. **Constructor:** Write `func NewTask(title string, priority int) Task` that
   returns a task with `Done: false` and an auto-incrementing ID (use a package
   variable as a counter).
5. **Embedding:** Make a `BaseEntity` struct with `ID int` and `CreatedAt string`,
   then embed it in both a `User` and a `Product` struct. Show that both get the
   `ID` and `CreatedAt` fields for free.

---

## ✅ Session 08 Checklist

- [ ] I can define a struct and create values 3 ways (named, positional, zero)
- [ ] I can access and modify struct fields
- [ ] I know uppercase = exported, lowercase = unexported
- [ ] I can write a method with a value receiver and a pointer receiver
- [ ] I know to use a pointer receiver when I need to modify the struct
- [ ] I follow the `NewX` constructor convention
- [ ] I understand struct embedding (composition over inheritance)
- [ ] I completed all 5 exercises

**Previous:** [← Session 07](session-07.md) · **Next:** [Session 09 — Pointers →](session-09.md)
