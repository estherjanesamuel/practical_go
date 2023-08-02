package isomorphicstrings

import (
	"fmt"
	"testing"
)

func TestIsomorphic(t *testing.T) {
	tests := []struct{
		s,t string
		want bool
	}{
		{"egg", "add",true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for _, ts := range tests {
		testIsomorphic := fmt.Sprintf("%v %v %v", ts.s,ts.t, ts.want)
		t.Run(testIsomorphic, func(t *testing.T) {
			got := Is_Isomorphic(ts.s, ts.t)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, !got)
			}
		})
	}
}
