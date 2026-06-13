> 🌐 **Language / زبان:** English (this file) · [فارسی](session-15.fa.md)

# Session 15 — Files, JSON & Encoding 📄

**Goal (1 hour):** Learn to move data in and out of your program. You'll read and
write **files**, and master **JSON** with `encoding/json` and **struct tags** —
the single most important serialization skill for backend work. Your REST API
speaks JSON, so this session is the bridge to the final project.

> **Recap from Session 14:** you toured the standard library. JSON and file I/O
> are two more standard-library packages you'll use in almost every real program.

---

## 1. JSON — the language of web APIs (25 min)

**JSON** (JavaScript Object Notation) is the universal format for sending data
between services. Go's `encoding/json` converts between Go values and JSON.

Two operations, opposite directions:

- **Marshal**: Go value → JSON (you produce JSON to send in a response).
- **Unmarshal**: JSON → Go value (you parse JSON from a request).

```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

// Marshal: struct -> JSON bytes
u := User{ID: 1, Name: "Sobhan", Email: "sobhan@example.com"}
data, err := json.Marshal(u)
// data = {"id":1,"name":"Sobhan","email":"sobhan@example.com"}

// Unmarshal: JSON bytes -> struct (pass a POINTER with &)
var parsed User
err = json.Unmarshal([]byte(`{"id":2,"name":"Ali"}`), &parsed)
```

> 🔑 **Two things to remember:**
> 1. `Unmarshal` needs a **pointer** (`&parsed`) so it can fill in your variable.
> 2. JSON only sees **exported** (uppercase) fields. A lowercase field is invisible
>    to `encoding/json`. (Capitalization matters again — Session 08!)

For human-readable output, use `json.MarshalIndent(v, "", "  ")`.

Run [`examples/session15/jsonbasics/jsonbasics.go`](../examples/session15/jsonbasics/jsonbasics.go).

---

## 2. Struct tags — controlling the JSON (20 min)

A **struct tag** is the backtick string after a field. For JSON it controls the
key name and behavior:

```go
type Product struct {
    ID       int     `json:"id"`                 // JSON key is "id", not "ID"
    Title    string  `json:"title"`
    Price    float64 `json:"price"`
    Discount float64 `json:"discount,omitempty"` // omit from JSON when zero
    Internal string  `json:"-"`                   // NEVER appears in JSON
}
```

| Tag | Effect |
|-----|--------|
| `json:"id"` | use `id` as the JSON key (lowercase, API-friendly) |
| `json:"discount,omitempty"` | leave the field out entirely when it's the zero value |
| `json:"-"` | never serialize this field (great for passwords, internal data) |

> 🔑 **`omitempty`** keeps your JSON clean by dropping empty/zero fields.
> **`json:"-"`** is how you keep secrets (like a password hash) out of API
> responses — you'll use it in the final project's `User` model.

When unmarshalling, unknown JSON keys are **ignored** and missing keys keep their
**zero value** — so your code is robust to extra or absent fields.

Run [`examples/session15/tags/tags.go`](../examples/session15/tags/tags.go).

---

## 3. Files — reading and writing (15 min)

For most cases, the whole-file helpers are all you need:

```go
// Write (creates or overwrites). 0644 = owner read/write, others read.
err := os.WriteFile("data.txt", []byte("hello\n"), 0644)

// Read the entire file into memory.
content, err := os.ReadFile("data.txt")   // content is []byte
fmt.Println(string(content))
```

Combine with JSON to persist structured data (a poor-man's database, and exactly
how config files work):

```go
cfg := Config{AppName: "TaskFlow", Port: 8080}
bytes, _ := json.MarshalIndent(cfg, "", "  ")
os.WriteFile("config.json", bytes, 0644)        // save

raw, _ := os.ReadFile("config.json")
var loaded Config
json.Unmarshal(raw, &loaded)                    // load
```

> 💡 **File permissions** like `0644` are octal Unix permissions: `6`=read+write
> for the owner, `4`=read for group and others. `0644` is the standard for a
> normal data file; `0755` (adds execute) is standard for directories/binaries.

> 📦 For **large** files you stream with `os.Open` + `bufio.Scanner` instead of
> loading everything into memory. We'll keep it simple here; `ReadFile`/`WriteFile`
> cover the vast majority of needs.

Run [`examples/session15/files/files.go`](../examples/session15/files/files.go).

---

## 🎯 Exercises (do these before Session 16!)

Create `examples/session15/practice/practice.go`:

1. **Round-trip:** Define a `Book` struct (`Title`, `Author`, `Year`, `Pages`)
   with JSON tags. Marshal one to pretty JSON, print it, then Unmarshal a JSON
   string back into a `Book` and print the struct.
2. **Hide a secret:** Add a `Password string` field to a `User` struct tagged
   `json:"-"`. Marshal it and confirm the password never appears in the output.
3. **omitempty:** Make a struct with an optional `Nickname string` using
   `omitempty`. Marshal it once empty and once filled; compare the JSON.
4. **Save & load a list:** Create a `[]Task`, marshal it, write it to a temp file,
   read it back, unmarshal it, and print the restored slice. Clean up the file.
5. **Config loader:** Write `func loadConfig(path string) (Config, error)` that
   reads a JSON file and returns a populated `Config` (or an error). Test it.

---

## ✅ Session 15 Checklist

- [ ] I can Marshal a Go value to JSON and Unmarshal JSON into a struct
- [ ] I remember Unmarshal needs a pointer (`&v`)
- [ ] I know only exported (uppercase) fields are seen by `encoding/json`
- [ ] I can rename JSON keys with struct tags
- [ ] I can use `omitempty` and `json:"-"`
- [ ] I can read and write whole files with `os.ReadFile`/`os.WriteFile`
- [ ] I can persist a struct to a file as JSON and load it back
- [ ] I completed all 5 exercises

**Previous:** [← Session 14](session-14.md) · **Next:** [Session 16 — Testing →](session-16.md)
