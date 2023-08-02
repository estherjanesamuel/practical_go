package mergesortedarray

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct{
		nums1,nums2 []int
		m,n int
		want []int
	} {
		{[]int{1,2,3,0,0,0},[]int{2,5,6}, 3,3,[]int{1,2,2,3,5,6}},
		{[]int{1},[]int{}, 1,0,[]int{1}},
		{[]int{0},[]int{1}, 0,1,[]int{1}},
	}

	for _, ts := range tests {
		testMergeSortedArray := fmt.Sprintf("mege %v %v %v %v %v", ts.nums1, ts.m,ts.nums2, ts.n, ts.want)
		t.Run(testMergeSortedArray, func(t *testing.T) {
			got := merge(ts.nums1, ts.m, ts.nums2, ts.n)
			if !slices.Equal(got, ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

