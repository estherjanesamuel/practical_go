package firstoccurrence

import (
	"fmt"
	"testing"
)

func TestFirstOccurrence(t *testing.T) {
	tests := []struct{
		heystack string
		needle string
		want int
	}{
		{"sadbutsad", "sad",0},
		{"leetcode", "leeto",-1},
	}

	for _, ts := range tests {
		testFirstOccurrence := fmt.Sprintf("%v %v %v",ts.heystack,ts.needle,ts.want)
		t.Run(testFirstOccurrence, func(t *testing.T) {
			got := firstOccurrence(ts.heystack, ts.needle)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
