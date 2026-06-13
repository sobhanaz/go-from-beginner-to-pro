> 🌐 **Language / زبان:** English (this file) · [فارسی](session-06.fa.md)

# Session 06 — Arrays & Slices 📚

**Goal (1 hour):** Learn how Go stores lists of values. You'll meet fixed-size
**arrays** (rarely used directly) and **slices** — the dynamic, growable list
that is the single most-used data structure in Go. By the end you'll `append`,
slice, iterate, and avoid the classic slice gotcha.

> **Recap from Session 05:** you've mastered functions (multiple returns, the
> `value, err` pattern, variadic, closures, `defer`). Now we hold *collections*
> of data.

---

## 1. Arrays — fixed size (10 min)

An **array** holds a fixed number of elements of the same type. The size is
part of the type and **cannot change**.

```go
var nums [3]int           // [0 0 0] — zero-valued
nums[0] = 10              // set by index (0-based)
nums[1] = 20
fmt.Println(nums)         // [10 20 0]
fmt.Println(len(nums))    // 3

// Declare with values:
primes := [5]int{2, 3, 5, 7, 11}

// Let Go count the size for you with ...
days := [...]string{"Mon", "Tue", "Wed"} // length 3
```

**Key facts:**
- Indexing is 0-based; `len()` gives the length.
- The size is fixed forever. `[3]int` and `[4]int` are *different types*.
- Arrays are **copied** when assigned or passed to functions (value semantics).

> 💡 In practice you'll rarely declare arrays directly — **slices** (next) are
> what you actually use. But arrays exist underneath every slice, so understand them.

---

## 2. Slices — the dynamic list you'll use everywhere (20 min)

A **slice** is a flexible, growable view into an array. No fixed size — it grows
as you add elements. This is *the* Go collection.

```go
// Literal — no size in the brackets = it's a slice, not an array.
fruits := []string{"apple", "banana"}

// Empty slice with make(): make([]T, length, capacity)
scores := make([]int, 0)      // empty, ready to grow
buffer := make([]int, 3)      // [0 0 0], length 3
```

### Growing a slice with `append`

`append` adds elements and returns a **new** slice. Always assign the result back:

```go
nums := []int{1, 2, 3}
nums = append(nums, 4)         // [1 2 3 4]
nums = append(nums, 5, 6, 7)   // append several at once
nums = append(nums, others...) // append another slice (spread with ...)
```

> 🔑 **Always write `nums = append(nums, x)`.** `append` may allocate a new
> underlying array, so it *returns* the updated slice. Ignoring the return is a
> classic beginner bug.

### Length vs capacity

A slice has a **length** (how many elements it holds) and a **capacity** (how
many it can hold before it must grow the underlying array):

```go
s := make([]int, 2, 5)
fmt.Println(len(s), cap(s))   // 2 5
```

You don't usually manage capacity by hand, but knowing it exists explains why
`append` is efficient (it grows in chunks, not one element at a time).

Run [`examples/session06/slices/slices.go`](../examples/session06/slices/slices.go).

---

## 3. Slicing — taking a sub-slice (15 min)

The `s[low:high]` syntax gives you a slice from index `low` up to **but not
including** `high`:

```go
nums := []int{10, 20, 30, 40, 50}
fmt.Println(nums[1:3])   // [20 30]   (indexes 1, 2)
fmt.Println(nums[:2])    // [10 20]   (from start)
fmt.Println(nums[3:])    // [40 50]   (to end)
fmt.Println(nums[:])     // [10 20 30 40 50] (whole thing)
```

Half-open ranges (`low` included, `high` excluded) mean `nums[1:3]` has exactly
`3 - 1 = 2` elements. This convention makes lengths easy to compute.

### ⚠️ The classic gotcha: slices share memory

A sub-slice points at the **same underlying array** as the original. Changing
one can change the other:

```go
original := []int{1, 2, 3, 4, 5}
part := original[1:3]   // [2 3], shares memory with original
part[0] = 99
fmt.Println(original)   // [1 99 3 4 5]  ← original changed too!
```

To get an independent copy, use `copy`:

```go
dst := make([]int, len(part))
copy(dst, part)         // dst is now a true, separate copy
```

This trips up *everyone* at first. Remember: slicing is a *view*, not a copy.

Run [`examples/session06/slicing/slicing.go`](../examples/session06/slicing/slicing.go).

---

## 4. Iterating with `for range` (10 min)

The natural way to walk a slice (you previewed this in Session 04):

```go
fruits := []string{"apple", "banana", "cherry"}

for i, fruit := range fruits {
    fmt.Printf("%d: %s\n", i, fruit)
}

// Only values? Discard the index with _:
for _, fruit := range fruits {
    fmt.Println(fruit)
}

// Only indexes? Just omit the second variable:
for i := range fruits {
    fmt.Println(i)
}
```

A common pattern — building a new slice from an old one (this is "map" in other
languages):

```go
prices := []float64{10, 20, 30}
withTax := make([]float64, 0, len(prices))
for _, p := range prices {
    withTax = append(withTax, p*1.09)
}
```

Run [`examples/session06/arrays/arrays.go`](../examples/session06/arrays/arrays.go) for the
array basics, then the others.

---

## 🎯 Exercises (do these before Session 07!)

Create `examples/session06/practice/practice.go`:

1. **Build a list:** Start with an empty `[]int`, `append` the numbers 1–10 in a
   loop, then print the slice and its length.
2. **Sum & average:** Given `nums := []float64{4, 8, 15, 16, 23, 42}`, loop with
   `range` to compute and print the sum and the average.
3. **Filter:** Write `func evens(nums []int) []int` that returns only the even
   numbers from the input slice. Test it.
4. **Reverse:** Write `func reverse(s []int) []int` that returns a new slice with
   the elements in reverse order.
5. **Prove the gotcha:** Create a slice, take a sub-slice with `[1:3]`, modify
   the sub-slice, and print the original to *see* the shared-memory effect. Then
   redo it with `copy` and show the original is now unaffected.

---

## ✅ Session 06 Checklist

- [ ] I know an array has a fixed size baked into its type
- [ ] I can create slices with a literal and with `make`
- [ ] I always write `s = append(s, x)` and know why
- [ ] I understand length vs capacity
- [ ] I can take sub-slices with `s[low:high]` (high is excluded)
- [ ] I understand that sub-slices share memory, and how `copy` avoids it
- [ ] I can iterate a slice three ways with `for range`
- [ ] I completed all 5 exercises

**Previous:** [← Session 05](session-05.md) · **Next:** [Session 07 — Maps, Strings & Runes →](session-07.md)
