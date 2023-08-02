package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T)  {
	tests := []struct {
		nums []int
		target int
		want int
	}{
		{[]int{4,5,6,7,0,1,2},0,4},
		{[]int{4,5,6,7,0,1,2},3,-1},
		{[]int{1},0,-1},
	}

	for _, ts := range tests {
		testSearch := fmt.Sprintf("%v %v %v", ts.nums, ts.target,ts.want)
		t.Run(testSearch, func(t *testing.T) {
			got := search(ts.nums, ts.target)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got,ts.want)
			}
		})
	}
}

