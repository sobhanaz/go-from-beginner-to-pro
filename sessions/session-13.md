> 🌐 **Language / زبان:** English (this file) · [فارسی](session-13.fa.md)

# Session 13 — Concurrency II: Patterns for Production ⚙️

**Goal (1 hour):** Move from "I can start a goroutine" to "I can write correct,
real-world concurrent code." You'll learn `sync.WaitGroup` (wait for many
goroutines), `sync.Mutex` (protect shared data), `select` (juggle multiple
channels), `context` (cancellation & timeouts), and the **worker pool** pattern
that ties it all together.

> **Recap from Session 12:** goroutines run concurrently; channels let them
> communicate and synchronize. Now we coordinate *many* of them correctly.

---

## 1. `sync.WaitGroup` — wait for many goroutines (15 min)

In Session 12 we used `time.Sleep` to wait — a guess. The real tool is
`sync.WaitGroup`, a counter that lets `main` block until all goroutines finish.

```go
var wg sync.WaitGroup

for i := 1; i <= 3; i++ {
    wg.Add(1)              // increment the counter before launching
    go func(id int) {
        defer wg.Done()    // decrement when this goroutine finishes
        fmt.Println("worker", id)
    }(i)
}

wg.Wait()                  // blocks until the counter hits 0
fmt.Println("all done")
```

Three calls, always in this shape:
- `wg.Add(1)` — "I'm about to start one goroutine." Call it *before* `go`.
- `defer wg.Done()` — "this goroutine is finished." First line inside the goroutine.
- `wg.Wait()` — "block here until all are done."

> ⚠️ **Pass the loop variable as an argument** (`go func(id int){...}(i)`) or, in
> Go 1.22+, rely on per-iteration loop variables. Capturing a shared loop variable
> by closure is a classic bug. Passing it explicitly is always safe.

Run [`examples/session13/waitgroup/waitgroup.go`](../examples/session13/waitgroup/waitgroup.go).

---

## 2. `sync.Mutex` — protect shared data (15 min)

When multiple goroutines write the **same** variable, you get a **data race** —
corrupted, unpredictable results. A `sync.Mutex` (mutual exclusion lock) ensures
only one goroutine touches the data at a time.

```go
type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.mu.Lock()         // acquire the lock
    defer c.mu.Unlock() // release it when the method returns
    c.value++           // safe: only one goroutine is here at a time
}
```

The `Lock`/`defer Unlock` pair is the standard pattern. Everything between is the
**critical section** — protected from concurrent access.

> 🔑 **Two ways to share state safely in Go:**
> 1. **Channels** — pass data *between* goroutines (preferred when it fits).
> 2. **Mutex** — guard data that goroutines genuinely *share* (counters, caches, maps).
>
> Use whichever is clearer. Run with `go run -race` to *prove* there's no race.

Run [`examples/session13/mutex/mutex.go`](../examples/session13/mutex/mutex.go) — try it with
and without `-race`, and try removing the lock to see the race detector fire.

---

## 3. `select` — wait on multiple channels (10 min)

`select` is like a `switch` for channels: it waits until *one* of several channel
operations is ready, then runs that case.

```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("timeout — nothing arrived in 1s")
default:
    fmt.Println("nothing ready right now (non-blocking)")
}
```

Key uses:
- **First-ready wins** — handle whichever channel produces a value first.
- **Timeouts** — `time.After(d)` returns a channel that fires after `d`; pair it
  with a `select` case to time out a wait.
- **Non-blocking** — a `default` case runs immediately if no channel is ready.

---

## 4. `context` — cancellation & deadlines (10 min)

`context.Context` is how Go propagates **cancellation**, **deadlines**, and
request-scoped values across goroutines and API calls. It's everywhere in real
Go — every HTTP handler and database call takes a `context`.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()   // always call cancel to release resources

select {
case result := <-doWork(ctx):
    fmt.Println("got:", result)
case <-ctx.Done():
    fmt.Println("cancelled or timed out:", ctx.Err())
}
```

A goroutine that respects context checks `ctx.Done()` (a channel that closes when
the context is cancelled or expires) and stops promptly. You'll thread `context`
through every layer of the final project's API.

> 🔑 **The pattern:** the top level creates a context (with a timeout or a cancel
> function), passes it down, and inner functions watch `ctx.Done()` to bail out
> early. Always `defer cancel()`.

Run [`examples/session13/selectctx/selectctx.go`](../examples/session13/selectctx/selectctx.go).

---

## 5. The worker pool — the pattern you'll actually use (10 min)

A **worker pool** runs a fixed number of goroutines that pull tasks off a "jobs"
channel and push results to a "results" channel. It limits concurrency (you don't
spawn a million goroutines) and is the workhorse of concurrent Go.

```go
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {        // pull jobs until the channel closes
        results <- j * j         // do the work, push the result
    }
}
```

`main` sends jobs, closes the jobs channel, waits for workers, then collects
results. This combines **everything**: goroutines, channels, directional types,
and `WaitGroup`. Study the example closely — it's a common interview question and
a building block for real systems.

Run [`examples/session13/workerpool/workerpool.go`](../examples/session13/workerpool/workerpool.go).

---

## 🎯 Exercises (do these before Session 14!)

Create `examples/session13/practice/practice.go` (run each with `-race`):

1. **WaitGroup:** Launch 5 goroutines that each print their ID. Use a `WaitGroup`
   so `main` waits for all of them — no `time.Sleep`.
2. **Safe counter:** Build a `Counter` struct with a `Mutex`. Launch 1000
   goroutines that each call `Inc()`. Print the final value — it must be exactly
   1000. Run with `-race` to confirm no race.
3. **Race demo:** Make the same counter *without* a mutex, run with `-race`, and
   observe the detector report the race (and a wrong total). Then fix it.
4. **Timeout with select:** Start a goroutine that sleeps a random time then sends
   on a channel. Use `select` with `time.After(500ms)` to print either the result
   or "timed out".
5. **Worker pool:** Build a pool of 3 workers that compute the square of jobs 1–9.
   Collect and print all 9 results.

---

## ✅ Session 13 Checklist

- [ ] I can use `WaitGroup` (Add/Done/Wait) to wait for many goroutines
- [ ] I pass loop variables into goroutines to avoid capture bugs
- [ ] I can protect shared data with `Mutex` (Lock / defer Unlock)
- [ ] I understand when to use channels vs a mutex
- [ ] I can use `select` for first-ready, timeouts, and non-blocking ops
- [ ] I understand `context` for cancellation/timeouts and always `defer cancel()`
- [ ] I can build a worker pool with jobs/results channels
- [ ] I run concurrent code with `-race`
- [ ] I completed all 5 exercises

**Previous:** [← Session 12](session-12.md) · **Next:** [Session 14 — Standard Library Tour →](session-14.md)

---

🎉 **Milestone:** Part 3 (The Go Way) complete! You now understand interfaces,
errors, and concurrency — the ideas that define idiomatic Go. From here it's all
about building real software: standard library, files/JSON, testing, and HTTP.
