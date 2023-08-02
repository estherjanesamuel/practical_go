package issubsequence

import (
	"fmt"
	"testing"
)

func TestIsSubsequence(t *testing.T) {

	tests := []struct{
		s,t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"aac", "ahagdc", true},
		{"axc", "ahbgdcx", false},
	}

	for _, ts := range tests {
		testIsSubsequence := fmt.Sprintf("%v %v %v", ts.s, ts.t, ts.want)
		t.Run(testIsSubsequence, func(t *testing.T) {
			got := Subsequence(ts.s, ts.t)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}

		})
	}


}

