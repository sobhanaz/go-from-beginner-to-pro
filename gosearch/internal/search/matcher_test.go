package search

import "testing"

func TestMatcher(t *testing.T) {
	cases := []struct {
		name       string
		pattern    string
		regex      bool
		ignoreCase bool
		line       string
		wantCount  int // number of matches expected
	}{
		{"literal hit", "go", false, false, "go gopher go", 3}, // also matches the "go" in "gopher"
		{"literal miss", "rust", false, false, "go gopher go", 0},
		{"case sensitive misses", "GO", false, false, "go gopher", 0},
		{"case insensitive hits", "GO", false, true, "Go gopher GOLANG", 3}, // Go, go(pher), GO(LANG)
		{"regex word boundary", `\bgo\b`, true, false, "go gopher go", 2},
		{"regex digits", `\d+`, true, false, "a1 b22 c333", 3},
		{"regex case-insensitive", `err`, true, true, "Error err ERR", 3},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m, err := NewMatcher(c.pattern, c.regex, c.ignoreCase)
			if err != nil {
				t.Fatalf("NewMatcher error: %v", err)
			}
			got := len(m.MatchLine(c.line))
			if got != c.wantCount {
				t.Errorf("MatchLine(%q) = %d matches; want %d", c.line, got, c.wantCount)
			}
		})
	}
}

func TestMatcherSpansAreCorrect(t *testing.T) {
	m, _ := NewMatcher("lo", false, false)
	spans := m.MatchLine("hello world, low")
	// "lo" appears at index 3 (hel-lo) and index 13 (low).
	want := [][2]int{{3, 5}, {13, 15}}
	if len(spans) != len(want) {
		t.Fatalf("got %v; want %v", spans, want)
	}
	for i := range want {
		if spans[i] != want[i] {
			t.Errorf("span %d = %v; want %v", i, spans[i], want[i])
		}
	}
}

func TestInvalidRegex(t *testing.T) {
	if _, err := NewMatcher("(unclosed", true, false); err == nil {
		t.Error("expected an error for an invalid regex, got nil")
	}
}
