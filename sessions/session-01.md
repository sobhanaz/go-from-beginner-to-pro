> 🌐 **Language / زبان:** English (this file) · [فارسی](session-01.fa.md)

# Session 01 — Hello, Go 🐹

**Goal (1 hour):** Understand what Go is, write and run your first program, and
learn the basic anatomy of every Go file. By the end you'll run real Go code on
your own machine.

---

## 1. What is Go? (5 min)

Go (also called **Golang**) is a programming language created at Google in 2009
by Robert Griesemer, Rob Pike, and Ken Thompson. It was designed to be:

- **Simple** — a small language you can learn fast (only ~25 keywords).
- **Fast** — compiles to a single native binary; runs close to C speed.
- **Built for concurrency** — doing many things at once is easy (goroutines).
- **Great for backends** — APIs, microservices, CLIs, DevOps tools (Docker,
  Kubernetes, and Terraform are all written in Go).

**Why learn it for a job?** Go developers are in high demand for backend and
cloud roles, the language is easy to be productive in, and the pay is strong.

---

## 2. The anatomy of a Go program (10 min)

Every Go program is made of **packages**. Here is the smallest complete program:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

Let's break down **every single line** — this matters, read carefully:

| Line | Meaning |
|------|---------|
| `package main` | Every Go file belongs to a package. `main` is special: it's the entry point of an executable program. |
| `import "fmt"` | Brings in the `fmt` ("format") package from the standard library — it handles printing. |
| `func main()` | A **function** named `main`. When you run the program, Go calls this first. It takes no arguments and returns nothing. |
| `fmt.Println(...)` | Calls the `Println` ("print line") function from the `fmt` package. It prints text and a newline. |
| `{ }` | Curly braces group the body of the function. In Go the opening `{` **must** be on the same line. |

> 🔑 **Rule:** a program that you can *run* must have `package main` and a
> `func main()`. That's the door Go walks through to start.

---

## 3. Run your first program (15 min)

The runnable file is at [`examples/session01/hello/hello.go`](../examples/session01/hello/hello.go).
Open a terminal in the course folder and run:

```bash
go run examples/session01/hello/hello.go
```

You should see:

```
Hello, Go!
```

🎉 You just compiled and ran a Go program! `go run` does two things at once:
it **compiles** your code to a temporary binary and **runs** it.

### `go run` vs `go build`

```bash
go run examples/session01/hello/hello.go     # compile + run, no file left behind (great for learning)

go build -o hello examples/session01/hello/hello.go   # creates a real binary called "hello"
./hello                                                # run that binary directly
```

The magic of Go: `go build` produces **one self-contained file** with no
dependencies. You can copy that single binary to another machine and it just runs.

---

## 4. Printing: the three you'll use most (10 min)

The `fmt` package has three printing functions you'll use all the time:

```go
fmt.Println("a", "b")        // prints: a b   (adds spaces + newline)
fmt.Print("no newline")      // prints exactly, no newline added
fmt.Printf("Name: %s\n", x)  // FORMATTED printing using "verbs"
```

`Printf` uses **format verbs** — placeholders that get replaced by values:

| Verb | Meaning | Example output |
|------|---------|----------------|
| `%s` | string | `Sobhan` |
| `%d` | integer (decimal) | `42` |
| `%f` | float | `3.140000` |
| `%t` | boolean | `true` |
| `%v` | **any value** (default format) | works for everything |
| `%T` | the **type** of the value | `int`, `string` |
| `\n` | newline | (line break) |

When in doubt, use `%v` — it prints any value sensibly. See it in action in
[`examples/session01/printing/printing.go`](../examples/session01/printing/printing.go):

```bash
go run examples/session01/printing/printing.go
```

---

## 5. Comments & formatting (5 min)

```go
// This is a single-line comment.

/*
   This is a
   multi-line comment.
*/
```

**Always format your code** with the official tool. It ends every style debate:

```bash
go fmt ./...
```

This auto-indents with tabs, aligns things, and fixes spacing. Real Go teams
require all code to be `gofmt`'d. Get in the habit now.

---

## 6. Common beginner errors (5 min)

Go is strict on purpose — these "errors" are teaching you good habits:

```go
import "fmt"   // ❌ error if you import something you don't use
```
> **"imported and not used"** — Go refuses to compile unused imports. Delete them.

```go
x := 5         // ❌ error if you declare a variable and never use it
```
> **"declared and not used"** — same idea for local variables. Use it or remove it.

These feel annoying at first but keep codebases clean. Embrace them.

---

## 🎯 Exercises (do these before Session 02!)

Create a new file `examples/session01/practice.go` and solve these:

1. **Your intro:** Print three lines — your name, your goal (e.g. "I want to be
   a Go developer"), and today's date — using `fmt.Println`.
2. **Use Printf:** Print the sentence `"I am learning Go and it is awesome"` but
   build it with `Printf` using `%s` for the words "Go" and "awesome".
3. **Types:** Use `%T` to print the type of `42`, `3.14`, `"hi"`, and `true`.
   Guess each one first, then check.
4. **Build it:** Run `go build -o myapp examples/session01/practice.go`, then
   run `./myapp`. Confirm you get the same output as `go run`.

> 💡 Stuck? The two solved examples in this folder show every technique you need.

---

## ✅ Session 01 Checklist

- [ ] I can explain what `package main` and `func main()` do
- [ ] I ran a program with `go run`
- [ ] I built a binary with `go build` and ran it
- [ ] I used `Println`, `Print`, and `Printf`
- [ ] I know what `%s`, `%d`, `%v`, and `%T` mean
- [ ] I ran `go fmt` on my code
- [ ] I completed all 4 exercises

**Next:** [Session 02 — Variables & Types →](session-02.md)
