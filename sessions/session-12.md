> 🌐 **Language / زبان:** English (this file) · [فارسی](session-12.fa.md)

# Session 12 — Concurrency I: Goroutines & Channels 🚦

**Goal (1 hour):** Meet the feature Go is famous for. A **goroutine** lets you
run a function *concurrently* (seemingly at the same time) with almost no effort,
and **channels** let goroutines communicate safely. By the end you'll launch
concurrent work and pass data between goroutines without locks.

> **Recap from Session 11:** you handle failures with `error` values. Concurrency
> introduces a new challenge — coordinating work that runs at the same time — and
> Go gives you elegant tools for it.

---

## 1. Concurrency vs. parallelism (5 min)

- **Concurrency** = *dealing with* many things at once (structure: tasks that can
  run independently and make progress in overlapping time).
- **Parallelism** = *doing* many things at the literal same instant (multiple CPU
  cores running code simultaneously).

Go makes **concurrency** easy to express; the Go runtime then schedules it across
cores to achieve parallelism when possible. You write concurrent code; Go handles
the hard parts.

---

## 2. Goroutines — concurrency with one keyword (15 min)

A **goroutine** is a function running independently and concurrently. Start one
by putting `go` before a function call:

```go
go doWork()          // runs doWork concurrently; main keeps going immediately
go fmt.Println("hi") // also works on any function call
```

That's it. Goroutines are extremely cheap — you can run *thousands*. They're not
OS threads; the Go runtime multiplexes many goroutines onto a few threads.

### The catch: main doesn't wait

When `main` returns, the program exits — **even if goroutines are still running.**
This naive version often prints nothing:

```go
func main() {
    go fmt.Println("from goroutine")
    // main ends immediately, goroutine may never get to run
}
```

You need a way to **wait** for goroutines to finish. The crude way is
`time.Sleep` (bad — you're guessing). The proper tools are `sync.WaitGroup`
(Session 13) and **channels** (next section), which also let goroutines *talk*.

Run [`examples/session12/goroutines/goroutines.go`](../examples/session12/goroutines/goroutines.go).

---

## 3. Channels — safe communication between goroutines (25 min)

A **channel** is a typed pipe: one goroutine sends values in, another receives
them out. Channels are how goroutines share data **safely** — no locks needed.

```go
ch := make(chan int)   // a channel that carries ints

ch <- 42               // SEND 42 into the channel (arrow points INTO ch)
value := <-ch          // RECEIVE from the channel (arrow points OUT of ch)
```

Remember the arrow direction: `ch <-` sends *into*, `<-ch` receives *out of*.

### Channels synchronize automatically

By default channels are **unbuffered**: a send blocks until another goroutine
receives, and a receive blocks until another goroutine sends. This blocking is a
*feature* — it synchronizes the two goroutines:

```go
func main() {
    ch := make(chan string)

    go func() {
        ch <- "result from goroutine"   // blocks until main receives
    }()

    msg := <-ch   // blocks until the goroutine sends; this is our "wait"!
    fmt.Println(msg)
}
```

Here the receive `<-ch` makes `main` wait for the goroutine — no `time.Sleep`
guessing. **"Don't communicate by sharing memory; share memory by
communicating"** is the Go concurrency motto.

### Buffered channels

A buffered channel holds a few values before blocking:

```go
ch := make(chan int, 3)   // capacity 3
ch <- 1
ch <- 2                   // these don't block until the buffer is full
```

### Closing and ranging over channels

The sender can `close` a channel to signal "no more values." A receiver can
`range` over a channel to read until it's closed:

```go
func produce(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)   // signal: done sending
}

ch := make(chan int)
go produce(ch)
for n := range ch {   // reads 1..5, then stops when ch is closed
    fmt.Println(n)
}
```

> 🔑 **Channel rules:**
> - Only the **sender** should `close` a channel, never the receiver.
> - Sending on a closed channel **panics**. Receiving from a closed channel
>   returns the zero value immediately (use comma-ok `v, ok := <-ch` to detect
>   closure: `ok` is `false` when closed and drained).

Run [`examples/session12/channels/channels.go`](../examples/session12/channels/channels.go).

---

## 4. Directional channels (10 min)

Functions can declare whether they only **send** or only **receive** on a
channel. This documents intent and lets the compiler catch mistakes:

```go
func producer(out chan<- int) {   // chan<- : send-only
    out <- 1
}

func consumer(in <-chan int) {    // <-chan : receive-only
    fmt.Println(<-in)
}
```

A plain `chan int` can do both; `chan<-` is send-only; `<-chan` is receive-only.
Using directional types in function signatures is a mark of clean Go.

Run [`examples/session12/directions/directions.go`](../examples/session12/directions/directions.go).

> ⚠️ **The race detector is your friend.** When goroutines touch shared data
> incorrectly, you get a *data race* (a subtle, nasty bug). Run your concurrent
> programs with `go run -race ...` to catch them. We'll lean on this in Session 13.

---

## 🎯 Exercises (do these before Session 13!)

Create `examples/session12/practice/practice.go`:

1. **Hello goroutine:** Launch a goroutine that prints a message, and use an
   unbuffered channel to make `main` wait for it before exiting.
2. **Sum via channel:** Write `func sum(nums []int, ch chan int)` that sends the
   total into a channel. Split a slice in half, sum each half in its own
   goroutine, then receive both partial sums and add them.
3. **Producer/consumer:** Write a producer that sends the numbers 1–10 into a
   channel and closes it, and a consumer (in `main`) that `range`s over the
   channel printing each. 
4. **Squares pipeline:** Make a goroutine that receives numbers on one channel and
   sends their squares on another. Feed it 1–5 and print the squares.
5. **Directional types:** Rewrite exercise 3 using `chan<- int` for the producer's
   parameter and `<-chan int` for the consumer's, so the compiler enforces direction.

> 💡 Run each with `go run -race examples/session12/practice/practice.go` to build
> the habit of checking for data races.

---

## ✅ Session 12 Checklist

- [ ] I can explain concurrency vs parallelism
- [ ] I can start a goroutine with the `go` keyword
- [ ] I understand that `main` exiting kills running goroutines
- [ ] I can create a channel and send (`ch <-`) / receive (`<-ch`)
- [ ] I know unbuffered channels block and thereby synchronize goroutines
- [ ] I can close a channel and `range` over it to receive until closed
- [ ] I know only the sender closes, and sending on a closed channel panics
- [ ] I can use directional channel types in function signatures
- [ ] I completed all 5 exercises

**Previous:** [← Session 11](session-11.md) · **Next:** [Session 13 — Concurrency II →](session-13.md)
