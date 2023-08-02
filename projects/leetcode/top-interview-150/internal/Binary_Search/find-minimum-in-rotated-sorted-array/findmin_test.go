package binarysearch

import (
	"fmt"
	"testing"
)

func TestFindMin(t *testing.T)  {
	tests := []struct{
		nums []int
		want int 
	} {
		{[]int{3,4,5,1,2}, 1},
		{[]int{4,5,6,7,0,1,2}, 0},
		{[]int{11,13,15,17}, 11},
	}

	for _, ts := range tests {
		testFindMin := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testFindMin, func(t *testing.T) {
			got := findMin(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

