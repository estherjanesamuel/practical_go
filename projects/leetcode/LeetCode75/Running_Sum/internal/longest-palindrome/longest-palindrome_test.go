package longestpalindrome

import (
	"fmt"
	"testing"
)

func TestLongestPalidrome(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abccccdd", 7},
	}

	for _, ts := range tests {
		testLongestPalidrome := fmt.Sprintf("%v %v ", ts.s, ts.want)
		t.Run(testLongestPalidrome, func(t *testing.T) {
			got := longestPalidrome(ts.s)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}

		})
	}

}
