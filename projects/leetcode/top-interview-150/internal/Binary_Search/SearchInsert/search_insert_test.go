package binarysearch

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T)  {
	tests := []struct{
		nums []int
		target int
		want int
	}{
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
		{[]int{}, -1, -1},
	}

	for _, ts := range tests {
		testSearchInsert := fmt.Sprintf("%v %v %v", ts.nums,ts.target, ts.want)
		t.Run(testSearchInsert, func(t *testing.T) {
			got := search_insert(ts.nums,ts.target)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

