package monotonicarray

import (
	"fmt"
	"testing"
)

func TestMonotonicArray(t *testing.T) {

	tests := []struct{
		nums []int
		want bool
	}{
		{[]int{1,2,2,3}, true},
		{[]int{6,5,4,4}, true},
		{[]int{1,3,2}, false},
	}

	for _, ts := range tests {
		testMonotonicArray := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testMonotonicArray, func(t *testing.T) {
			got := isMonotonic(ts.nums)

			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

