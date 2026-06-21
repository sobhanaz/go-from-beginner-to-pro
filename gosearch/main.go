// gosearch — a fast, concurrent text search for the command line (a mini-grep).
//
// Usage:
//
//	gosearch [flags] <pattern> [path]
//
// Examples:
//
//	gosearch func .                     # find "func" under the current dir
//	gosearch -i -ext .go "todo" ./src   # case-insensitive, only .go files
//	gosearch -r "func \w+\(" .          # regex search
//	gosearch -c "import" .              # count matches per file
package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"

	"gosearch/internal/search"
)

// ANSI colors (only used when writing to a real terminal).
const (
	cReset = "\033[0m"
	cPath  = "\033[36m"   // cyan
	cLine  = "\033[32m"   // green
	cMatch = "\033[1;31m" // bold red
)

func main() {
	ignoreCase := flag.Bool("i", false, "case-insensitive match")
	useRegex := flag.Bool("r", false, "treat the pattern as a regular expression")
	ext := flag.String("ext", "", "only search files with this extension, e.g. .go")
	workers := flag.Int("w", runtime.NumCPU(), "number of concurrent workers")
	hidden := flag.Bool("hidden", false, "include hidden files and directories")
	countOnly := flag.Bool("c", false, "print only a count of matches per file")
	noColor := flag.Bool("no-color", false, "disable colored output")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "gosearch — a fast concurrent text search.")
		fmt.Fprintln(os.Stderr, "\nUsage:\n  gosearch [flags] <pattern> [path]")
		fmt.Fprintln(os.Stderr, "\nFlags:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExamples:")
		fmt.Fprintln(os.Stderr, "  gosearch func .")
		fmt.Fprintln(os.Stderr, "  gosearch -i -ext .go \"todo\" ./src")
		fmt.Fprintln(os.Stderr, "  gosearch -r \"func \\w+\\(\" .")
	}
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}
	pattern := args[0]
	root := "."
	if len(args) >= 2 {
		root = args[1]
	}

	s, err := search.New(search.Options{
		Pattern:    pattern,
		Root:       root,
		IgnoreCase: *ignoreCase,
		Regex:      *useRegex,
		Ext:        *ext,
		Workers:    *workers,
		Hidden:     *hidden,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(2)
	}

	matches, err := s.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "warning:", err)
	}

	color := !*noColor && isTerminal(os.Stdout)
	if *countOnly {
		printCounts(matches, color)
	} else {
		printMatches(matches, color)
	}

	// Follow grep's convention: exit 1 when nothing matched.
	if len(matches) == 0 {
		os.Exit(1)
	}
}

// printMatches prints one line per hit: path:line: text (matches highlighted).
func printMatches(matches []search.Match, color bool) {
	for _, m := range matches {
		if color {
			fmt.Printf("%s%s%s:%s%d%s: %s\n",
				cPath, m.Path, cReset, cLine, m.Line, cReset,
				highlight(m.Text, m.Spans))
		} else {
			fmt.Printf("%s:%d: %s\n", m.Path, m.Line, m.Text)
		}
	}
}

// printCounts prints "path: N" for each file that had matches.
func printCounts(matches []search.Match, color bool) {
	counts := map[string]int{}
	order := []string{}
	for _, m := range matches {
		if counts[m.Path] == 0 {
			order = append(order, m.Path)
		}
		counts[m.Path]++
	}
	for _, p := range order {
		if color {
			fmt.Printf("%s%s%s: %d\n", cPath, p, cReset, counts[p])
		} else {
			fmt.Printf("%s: %d\n", p, counts[p])
		}
	}
}

// highlight wraps each matched span in red ANSI codes.
func highlight(text string, spans [][2]int) string {
	var b strings.Builder
	prev := 0
	for _, sp := range spans {
		if sp[0] < prev || sp[1] > len(text) {
			continue // skip overlapping/out-of-range spans defensively
		}
		b.WriteString(text[prev:sp[0]])
		b.WriteString(cMatch)
		b.WriteString(text[sp[0]:sp[1]])
		b.WriteString(cReset)
		prev = sp[1]
	}
	b.WriteString(text[prev:])
	return b.String()
}

// isTerminal reports whether f is a real terminal (so we only colorize there,
// not when piping into a file or another program).
func isTerminal(f *os.File) bool {
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}
