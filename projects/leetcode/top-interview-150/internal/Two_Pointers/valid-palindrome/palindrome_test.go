package validpalindrome

import (
	"fmt"
	"testing"
)

func TestValidPalindrome(t *testing.T) {
	tests := []struct{
		s string
		want bool
	}{
		{"A man, a plan, a canal: Panama",true},
		{"race a car",false},
		{"123456789",false},
	}

	for _, ts := range tests {
		testValidPalindrome := fmt.Sprintf("%v %v", ts.s, ts.want)
		t.Run(testValidPalindrome, func(t *testing.T) {
			got := isPalindrome(ts.s)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
