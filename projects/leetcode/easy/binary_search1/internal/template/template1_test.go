package template

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums []int
		target, want int
	}{
		{[]int{-1,0,3,5,9,12},9,4},
		{[]int{-1,0,3,5,9,12,14,17,18,21,25},14,6},
		{[]int{-1,0,3,5,9,12,14,17,18,21,25,27,31,35,37,40,42,46,49,52},14,6},
		{[]int{-1,0,3,5,9,12},1,-1},
	}
	for _, ts := range tests {
		binarySearch := fmt.Sprintf("%v %v", ts.nums,ts.target)
		t.Run(binarySearch, func(t *testing.T) {
			got := binarySeacrh(ts.nums,ts.target)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
	
}