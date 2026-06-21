package search

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
)

// Options configures a search run.
type Options struct {
	Pattern    string
	Root       string // directory (or file) to search
	IgnoreCase bool
	Regex      bool
	Ext        string // only files with this extension, e.g. ".go" ("" = any)
	Workers    int    // number of concurrent workers (0 = NumCPU)
	Hidden     bool   // include hidden files/dirs (those starting with ".")
}

// Match is a single matching line within a file.
type Match struct {
	Path  string
	Line  int
	Text  string
	Spans [][2]int // byte ranges of the matches within Text (for highlighting)
}

// Searcher runs a concurrent search defined by Options.
type Searcher struct {
	opts    Options
	matcher Matcher
}

// New validates the options, compiles the matcher, and returns a Searcher.
func New(opts Options) (*Searcher, error) {
	if opts.Workers <= 0 {
		opts.Workers = runtime.NumCPU()
	}
	if opts.Root == "" {
		opts.Root = "."
	}
	m, err := NewMatcher(opts.Pattern, opts.Regex, opts.IgnoreCase)
	if err != nil {
		return nil, err
	}
	return &Searcher{opts: opts, matcher: m}, nil
}

// Run walks the directory tree and returns all matches, sorted by path+line.
//
// It's a worker pool (Session 13): one goroutine walks the tree feeding file
// paths into `jobs`; N workers read files concurrently and send hits to
// `results`; a closer goroutine closes `results` once all workers finish.
func (s *Searcher) Run() ([]Match, error) {
	jobs := make(chan string, 64)
	results := make(chan Match, 64)
	var wg sync.WaitGroup

	// Start the workers.
	for i := 0; i < s.opts.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range jobs {
				s.searchFile(path, results)
			}
		}()
	}

	// Walk the tree in its own goroutine, feeding the jobs channel.
	walkErrCh := make(chan error, 1)
	go func() {
		err := filepath.WalkDir(s.opts.Root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil // skip entries we can't read, keep going
			}
			name := d.Name()
			// Skip hidden files/dirs unless asked to include them.
			if !s.opts.Hidden && path != s.opts.Root && strings.HasPrefix(name, ".") {
				if d.IsDir() {
					return filepath.SkipDir
				}
				return nil
			}
			if d.IsDir() {
				return nil
			}
			if s.opts.Ext != "" && filepath.Ext(path) != s.opts.Ext {
				return nil
			}
			jobs <- path
			return nil
		})
		close(jobs) // no more files -> workers' range loops will end
		walkErrCh <- err
	}()

	// Close results once every worker has finished.
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect all matches.
	var matches []Match
	for m := range results {
		matches = append(matches, m)
	}

	// Deterministic output order.
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Path != matches[j].Path {
			return matches[i].Path < matches[j].Path
		}
		return matches[i].Line < matches[j].Line
	})

	return matches, <-walkErrCh
}

// searchFile scans a single file line by line and sends any matches out.
func (s *Searcher) searchFile(path string, out chan<- Match) {
	f, err := os.Open(path)
	if err != nil {
		return // unreadable file: skip silently
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024) // tolerate long lines

	lineNo := 0
	for sc.Scan() {
		lineNo++
		line := sc.Text()
		// A NUL byte strongly suggests a binary file — skip the rest of it.
		if strings.IndexByte(line, 0) >= 0 {
			return
		}
		if spans := s.matcher.MatchLine(line); len(spans) > 0 {
			out <- Match{Path: path, Line: lineNo, Text: line, Spans: spans}
		}
	}
}
