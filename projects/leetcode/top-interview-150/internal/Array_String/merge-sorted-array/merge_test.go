package mergesortedarray

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestMerge(t *testing.T) {
	tests := []struct{
		arr,nums []int
		m, n int
		want []int
	}{
		{[]int{1,2,5,0,0,0},[]int{2,3,4},3,3,[]int{1,2,2,3,4,5}},
		{[]int{0},[]int{1},0,1,[]int{1}},
	}

	for _, ts := range tests {
		testMerge := fmt.Sprintf("%v %v %v %v", ts.arr,ts.nums,ts.m,ts.n)
		t.Run(testMerge, func(t *testing.T) {
			got := merge(ts.arr, ts.m, ts.nums, ts.n)
			if !slices.Equal(got, ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

