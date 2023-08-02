package firstoccurrenceinastring

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct{
		heystack, needle string
		want int
	} {
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
	}
	for _, ts := range tests {
		testStrStr := fmt.Sprintf("strStr %v %v %v", ts.heystack,ts.needle, ts.want)
		t.Run(testStrStr, func(t *testing.T) {
			got := strStr(ts.heystack, ts.needle)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
		testStrStrII := fmt.Sprintf("strStrII %v %v %v", ts.heystack,ts.needle, ts.want)
		t.Run(testStrStrII, func(t *testing.T) {
			got := strStrII(ts.heystack, ts.needle)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
		}
}

