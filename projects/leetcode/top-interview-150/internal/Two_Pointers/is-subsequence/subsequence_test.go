package issubsequence

import (
	"fmt"
	"testing"
)

func TestSubsequence(t *testing.T) {
	tests := []struct{
		s,t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"axc", "ahgdcx", false},
		{"b", "abc", true},
		{"cb", "abc", false},
		{"", "ahgdcx", true},
	}
	for _, ts := range tests {
		testSubsequence := fmt.Sprintf("%v %v %v", ts.s,ts.t,ts.want)
		t.Run(testSubsequence, func(t *testing.T) {
			got := isSubsequence(ts.s, ts.t)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
