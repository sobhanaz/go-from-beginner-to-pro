package search

import (
	"os"
	"path/filepath"
	"testing"
)

// buildTree creates a temporary directory tree to search over.
func buildTree(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	files := map[string]string{
		"a.txt":          "hello world\nTODO: fix this\nsecond line\n",
		"b.go":           "package main\n// TODO add tests\nfunc main() {}\n",
		"sub/c.go":       "package sub\nfunc Helper() {}\n",
		"sub/notes.md":   "a todo here\nand a TODO there\n",
		".hidden/secret": "TODO hidden\n",
	}
	for name, content := range files {
		full := filepath.Join(root, name)
		if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func runSearch(t *testing.T, opts Options) []Match {
	t.Helper()
	s, err := New(opts)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	matches, err := s.Run()
	if err != nil {
		t.Fatalf("Run: %v", err)
	}
	return matches
}

func TestSearchLiteral(t *testing.T) {
	root := buildTree(t)
	// "TODO" (case-sensitive) appears in a.txt, b.go, sub/notes.md — but NOT in
	// the hidden file (hidden is skipped by default).
	matches := runSearch(t, Options{Pattern: "TODO", Root: root})
	if len(matches) != 3 {
		t.Fatalf("got %d matches; want 3:\n%+v", len(matches), matches)
	}
}

func TestSearchCaseInsensitive(t *testing.T) {
	root := buildTree(t)
	// "todo" case-insensitive: a.txt(1) + b.go(1) + notes.md(2) = 4
	matches := runSearch(t, Options{Pattern: "todo", Root: root, IgnoreCase: true})
	if len(matches) != 4 {
		t.Fatalf("got %d; want 4:\n%+v", len(matches), matches)
	}
}

func TestSearchExtensionFilter(t *testing.T) {
	root := buildTree(t)
	// Only .go files: b.go has "func main", sub/c.go has "func Helper".
	matches := runSearch(t, Options{Pattern: "func", Root: root, Ext: ".go"})
	if len(matches) != 2 {
		t.Fatalf("got %d; want 2:\n%+v", len(matches), matches)
	}
	for _, m := range matches {
		if filepath.Ext(m.Path) != ".go" {
			t.Errorf("matched a non-.go file: %s", m.Path)
		}
	}
}

func TestSearchHiddenIncluded(t *testing.T) {
	root := buildTree(t)
	withHidden := runSearch(t, Options{Pattern: "TODO", Root: root, Hidden: true})
	if len(withHidden) != 4 { // now the .hidden/secret file counts too
		t.Fatalf("with hidden: got %d; want 4", len(withHidden))
	}
}

func TestSearchRegex(t *testing.T) {
	root := buildTree(t)
	// Lines that start with "func ".
	matches := runSearch(t, Options{Pattern: `^func `, Root: root, Regex: true})
	if len(matches) != 2 {
		t.Fatalf("got %d; want 2:\n%+v", len(matches), matches)
	}
}

func TestSearchNoMatch(t *testing.T) {
	root := buildTree(t)
	matches := runSearch(t, Options{Pattern: "zzz-not-here", Root: root})
	if len(matches) != 0 {
		t.Fatalf("got %d; want 0", len(matches))
	}
}

// Results must be deterministic (sorted by path then line) regardless of how
// many concurrent workers run.
func TestSearchDeterministicOrder(t *testing.T) {
	root := buildTree(t)
	for _, workers := range []int{1, 4, 16} {
		matches := runSearch(t, Options{Pattern: "e", Root: root, Workers: workers})
		for i := 1; i < len(matches); i++ {
			a, b := matches[i-1], matches[i]
			if a.Path > b.Path || (a.Path == b.Path && a.Line > b.Line) {
				t.Fatalf("workers=%d: results not sorted at %d: %v then %v",
					workers, i, a, b)
			}
		}
	}
}
