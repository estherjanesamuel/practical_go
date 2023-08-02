package binarysearch

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T)  {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{1,2,3,1}, 2},
		{[]int{1,2,1,3,5,6,4}, 5},
	}

	for _, ts := range tests {
		testFindPeakElement := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testFindPeakElement, func(t *testing.T) {
			got := findPeakElement(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}


