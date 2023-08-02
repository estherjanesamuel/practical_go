package wordpattern

import (
	"fmt"
	"testing"
)

func TestWordPatter(t *testing.T) {
	tests := []struct {
		pattern, s string
		want bool
	}{
		// {"abba", "dog cat cat dog", true},
		{"abba", "dog cat cat fish", false},
		// {"aaaa", "dog cat cat dog", false},
		{"aaa","aa aa aa aa", false},
	}
	for _, ts := range tests {
		testWordPattern := fmt.Sprintf("%v %v %v", ts.pattern, ts.s, ts.want)
		t.Run(testWordPattern, func(t *testing.T) {
			got := samePattern(ts.pattern, ts.s)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got,ts.want)
			}
		})
	}
}

