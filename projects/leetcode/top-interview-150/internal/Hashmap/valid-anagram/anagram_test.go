package validanagram

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T)  {
	tests := []struct{
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, ts := range tests {
		testValidAnagram := fmt.Sprintf("%v %v %v", ts.s, ts.t, ts.want)
		t.Run(testValidAnagram, func(t *testing.T) {
			got := isValidAnagram(ts.s, ts.t)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
