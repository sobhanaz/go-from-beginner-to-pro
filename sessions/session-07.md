> 🌐 **Language / زبان:** English (this file) · [فارسی](session-07.fa.md)

# Session 07 — Maps, Strings & Runes 🗺️

**Goal (1 hour):** Learn Go's key/value store (the **map**), and finally
understand what a **string** really is under the hood — including the difference
between a `byte` and a `rune`, and why that matters the moment you touch
non-English text or emoji.

> **Recap from Session 06:** slices are ordered, growable lists. Maps are the
> *unordered* lookup-by-key counterpart. Together they cover most data you'll model.

---

## 1. Maps — key/value lookup (20 min)

A **map** associates **keys** with **values**. Think of a dictionary: look up a
word (key) to get its definition (value). Lookups are fast (roughly O(1)).

```go
// map[KeyType]ValueType
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

fmt.Println(ages["Alice"])   // 30
ages["Carol"] = 28           // add or update
ages["Alice"] = 31           // update existing
delete(ages, "Bob")          // remove a key
fmt.Println(len(ages))       // number of entries
```

### Creating maps

```go
m1 := map[string]int{}          // empty literal
m2 := make(map[string]int)      // same, with make
var m3 map[string]int           // nil map — can READ but NOT write (panics!)
```

> ⚠️ A `nil` map (declared with `var` and no value) can be read from but
> **panics if you write to it**. Always create maps with a literal or `make`
> before adding keys.

### The "comma ok" idiom — does a key exist?

Reading a missing key returns the **zero value**, so you can't tell "missing"
from "present but zero". The fix is the two-value form:

```go
age, ok := ages["Dave"]
if ok {
    fmt.Println("Dave is", age)
} else {
    fmt.Println("No Dave in the map")
}
```

`ok` is `true` only if the key actually exists. This **comma-ok** pattern is
idiomatic Go — you'll see it constantly.

### Iterating a map

```go
for name, age := range ages {
    fmt.Printf("%s is %d\n", name, age)
}
```

> 🔑 **Map order is random!** Go deliberately randomizes iteration order so you
> never accidentally depend on it. If you need sorted output, collect the keys
> into a slice and sort them (we'll use `sort` in Session 14).

Run [`examples/session07/maps/maps.go`](../examples/session07/maps/maps.go).

---

## 2. Strings — immutable bytes (15 min)

A Go `string` is a **read-only sequence of bytes**. Two important truths:

1. **Strings are immutable** — you can't change a character in place. To "modify"
   a string you build a new one.
2. **A string is UTF-8 encoded bytes.** For plain English, 1 character = 1 byte.
   For other scripts (Persian, Chinese) and emoji, **one character can be
   several bytes**.

```go
s := "Hello"
fmt.Println(len(s))      // 5  (number of BYTES)
fmt.Println(s[0])        // 72 (the BYTE value of 'H', not "H"!)
fmt.Printf("%c\n", s[0]) // H  (%c formats a byte/rune as a character)
```

> ⚠️ `len(s)` counts **bytes**, not characters. `len("héllo")` is 6, not 5,
> because `é` is 2 bytes in UTF-8. Same for `len("سلام")` — it's bytes, not letters.

### Common string operations

The `strings` package (full tour in Session 14) has the everyday helpers:

```go
import "strings"

strings.ToUpper("hello")            // "HELLO"
strings.Contains("hello", "ell")    // true
strings.Split("a,b,c", ",")         // ["a" "b" "c"]
strings.Replace("aaa", "a", "b", -1)// "bbb"
strings.TrimSpace("  hi  ")         // "hi"

// Joining strings efficiently:
strings.Join([]string{"a", "b"}, "-") // "a-b"
```

Run [`examples/session07/strings/strings.go`](../examples/session07/strings/strings.go).

---

## 3. Bytes vs. Runes — handling any language (20 min)

This is the part that confuses everyone, so go slow.

- A **`byte`** is a single 8-bit value (`uint8`). A string is made of bytes.
- A **`rune`** is a single Unicode *character* (`int32`). It's Go's way of saying
  "one logical character, no matter how many bytes it takes."

When you need to work with **characters** (not bytes), convert to `[]rune` or
range over the string (which yields runes):

```go
word := "héllo"

// Ranging over a string gives RUNES with their byte index.
for i, r := range word {
    fmt.Printf("byte %d: %c (rune %d)\n", i, r, r)
}

// Count CHARACTERS correctly:
fmt.Println(len(word))            // 6  (bytes)
fmt.Println(len([]rune(word)))    // 5  (characters)

// Or the standard helper:
import "unicode/utf8"
fmt.Println(utf8.RuneCountInString(word)) // 5
```

> 🔑 **Rule of thumb:**
> - Working with raw data / ASCII? Bytes are fine.
> - Counting or indexing real characters (Persian, emoji, accents)? Use **runes**
>   (`[]rune(s)` or `range`), never `s[i]`.

### Building strings efficiently

Because strings are immutable, repeatedly doing `s = s + x` in a loop is wasteful
(it allocates a new string every time). For lots of concatenation, use
`strings.Builder`:

```go
var b strings.Builder
for i := 0; i < 5; i++ {
    b.WriteString("go ")
}
result := b.String()   // "go go go go go "
```

Run [`examples/session07/runes/runes.go`](../examples/session07/runes/runes.go).

---

## 🎯 Exercises (do these before Session 08!)

Create `examples/session07/practice/practice.go`:

1. **Word counter:** Given a sentence string, split it on spaces and use a
   `map[string]int` to count how many times each word appears. Print the map.
2. **Comma-ok:** Build a `map[string]float64` of product prices. Look up a product
   that exists and one that doesn't, using the comma-ok idiom to report clearly.
3. **Vowel counter:** Write `func countVowels(s string) int` that counts vowels,
   ranging over the string (runes), case-insensitive.
4. **Reverse a string correctly:** Write `func reverse(s string) string` that
   reverses by `[]rune` (not bytes) so it works on `"héllo"` and emoji. Test it.
5. **Char vs byte length:** Print both `len(s)` and `len([]rune(s))` for a string
   containing a non-ASCII character, and explain the difference in a comment.

---

## ✅ Session 07 Checklist

- [ ] I can create, read, update, and delete entries in a map
- [ ] I know a nil map panics on write, so I create maps with `make`/literal
- [ ] I use the comma-ok idiom to check if a key exists
- [ ] I know map iteration order is random
- [ ] I understand a string is immutable UTF-8 bytes, and `len` counts bytes
- [ ] I know the difference between a `byte` and a `rune`
- [ ] I can correctly count/iterate characters using runes
- [ ] I completed all 5 exercises

**Previous:** [← Session 06](session-06.md) · **Next:** [Session 08 — Structs & Methods →](session-08.md)
