package math

import (
	"fmt"
	"testing"
)

func TestPalindrome(t *testing.T)  {
	tests := []struct{
		x int
		want bool
	}{
		// {121, true},
		// {-121, false},
		{10, false},
		// {101, false},
		// {121121, true},
	}

	for _, ts := range tests {
		testPalindrome := fmt.Sprintf("%v %v", ts.x,ts.want)
		t.Run(testPalindrome, func(t *testing.T) {
			got := is_palindrome(ts.x)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

