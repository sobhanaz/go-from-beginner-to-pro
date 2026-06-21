// Package search implements a fast, concurrent text search over a directory.
package search

import (
	"fmt"
	"regexp"
	"strings"
)

// Matcher decides whether a line matches the pattern and WHERE, so callers can
// highlight the matches. It's an interface (Session 10) so we can swap a plain
// substring matcher for a regex matcher without changing the search engine.
type Matcher interface {
	// MatchLine returns the [start,end) byte ranges of every match in line.
	// An empty slice means the line did not match.
	MatchLine(line string) [][2]int
}

// literalMatcher matches a fixed, case-sensitive substring. This is the fast path.
type literalMatcher struct {
	pattern string
}

func (m literalMatcher) MatchLine(line string) [][2]int {
	var spans [][2]int
	from := 0
	for {
		i := strings.Index(line[from:], m.pattern)
		if i < 0 {
			break
		}
		start := from + i
		spans = append(spans, [2]int{start, start + len(m.pattern)})
		from = start + len(m.pattern)
	}
	return spans
}

// regexMatcher matches a regular expression (also used for case-insensitive
// literal search, via a quoted pattern, so highlight offsets stay correct).
type regexMatcher struct {
	re *regexp.Regexp
}

func (m regexMatcher) MatchLine(line string) [][2]int {
	locs := m.re.FindAllStringIndex(line, -1)
	spans := make([][2]int, 0, len(locs))
	for _, l := range locs {
		spans = append(spans, [2]int{l[0], l[1]})
	}
	return spans
}

// NewMatcher builds the right Matcher for the given options.
//   - regex            -> compile the pattern as a regular expression
//   - literal + fold   -> compile a case-insensitive regex of the quoted pattern
//   - literal          -> fast substring matcher
func NewMatcher(pattern string, useRegex, ignoreCase bool) (Matcher, error) {
	if useRegex {
		expr := pattern
		if ignoreCase {
			expr = "(?i)" + expr
		}
		re, err := regexp.Compile(expr)
		if err != nil {
			return nil, fmt.Errorf("invalid regex %q: %w", pattern, err)
		}
		return regexMatcher{re: re}, nil
	}
	if ignoreCase {
		re := regexp.MustCompile("(?i)" + regexp.QuoteMeta(pattern))
		return regexMatcher{re: re}, nil
	}
	return literalMatcher{pattern: pattern}, nil
}
