# gosearch 🔎

A fast, **concurrent** command-line text search — a mini `grep`/`ripgrep`, written
in Go with **only the standard library**. It recursively searches a directory,
matching plain text or regular expressions, using a **worker pool** so large
trees are scanned in parallel.

> Built to showcase Go's concurrency model (goroutines, channels, worker pools)
> in a small, genuinely useful CLI tool. No third-party dependencies.

![gosearch in action](docs/demo.png)

---

## ✨ Features

- 🔁 **Concurrent** directory walk + file scanning via a worker pool (one goroutine per CPU by default)
- 🔤 **Plain** substring search (fast path) or full **regular expressions** (`-r`)
- 🔡 **Case-insensitive** matching (`-i`)
- 📁 **Extension filter** (`-ext .go`) and hidden-file control (`-hidden`)
- 🎨 **Colored output** with matched text highlighted (auto-disabled when piped)
- #️⃣ **Count mode** (`-c`) — matches per file
- 🧪 Tested (table-driven + temp-dir integration tests), **race-clean**, ~93% coverage
- 📦 Compiles to a single static binary — no runtime dependencies

## 🚀 Install & build

```bash
cd gosearch
go build -o gosearch .      # produces ./gosearch
# optional: install into your $GOBIN
go install .
```

## 📖 Usage

```
gosearch [flags] <pattern> [path]
```

`path` defaults to the current directory (`.`).

| Flag | Meaning | Default |
|------|---------|---------|
| `-i` | case-insensitive match | off |
| `-r` | treat the pattern as a regular expression | off |
| `-ext` | only search files with this extension (e.g. `.go`) | all files |
| `-w` | number of concurrent workers | number of CPUs |
| `-hidden` | include hidden files and directories | off |
| `-c` | print only a count of matches per file | off |
| `-no-color` | disable colored output | off |

## 💡 Examples

```bash
# every line containing "func" under the current directory
gosearch func .

# case-insensitive "todo", only in .go files, under ./src
gosearch -i -ext .go "todo" ./src

# regex: lines that are a function definition
gosearch -r "^func \w+\(" .

# how many times "import" appears per file
gosearch -c import .
```

Example output (matches are highlighted in red in a real terminal):

```
internal/search/search.go:40: func New(opts Options) (*Searcher, error) {
internal/search/search.go:59: func (s *Searcher) Run() ([]Match, error) {
main.go:33: func main() {
```

## 🏗️ How it works

```
main.go                      CLI: parse flags, print results (with color)
internal/search/
├── matcher.go               Matcher interface: literal (fast) + regex impls
├── search.go                Searcher: concurrent walk + worker pool
├── matcher_test.go          table-driven matcher tests
└── search_test.go           integration tests over a temp directory tree
```

The search is a classic **worker pool** (the pattern from the concurrency
chapter):

1. One goroutine walks the tree with `filepath.WalkDir`, pushing file paths into
   a `jobs` channel (skipping hidden/binary files and non-matching extensions).
2. `N` worker goroutines read from `jobs`, scan each file line-by-line, and push
   hits into a `results` channel.
3. A `sync.WaitGroup` tracks the workers; when they're all done, a closer
   goroutine closes `results`.
4. `main` drains `results` and sorts them for **deterministic** output.

Matching is behind a small **`Matcher` interface**, so the engine doesn't care
whether you're doing literal or regex search — a clean example of Go's implicit
interfaces.

> ⚙️ The whole thing is bounded by the worker count (`-w`), so it stays fast on
> huge trees without spawning a goroutine per file.

## ✅ Test

```bash
go test ./...                # all tests
go test -race ./...          # verify there are no data races
go test -cover ./...         # coverage (~93%)
```

## 🧱 Tech

- Go standard library only: `flag`, `bufio`, `os`, `io/fs`, `path/filepath`,
  `regexp`, `sync`, `sort`, `runtime`
- Concurrency: goroutines, buffered channels, `sync.WaitGroup`, worker pool

## 📝 License

MIT — built for learning and portfolio use.
