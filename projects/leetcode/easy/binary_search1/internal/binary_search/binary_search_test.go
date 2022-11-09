package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct{
		arr []int
		target, want int
	}{
		{[]int{-1,0,3,5,9,12}, 9, 4},
		{[]int{-1,0,3,5,9,12}, 2, -1},
		{[]int{-1,0,3,5,9,12}, 12, 5},
		{[]int{-1,0,3,5,9,12}, 5, 3},
	}

	for _, ts := range tests {
		testBinarySearch := fmt.Sprintf("%v, %v", ts.arr, ts.target)
		t.Run(testBinarySearch, func(t *testing.T) {
			got:= binarySearch(ts.arr, ts.target)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}