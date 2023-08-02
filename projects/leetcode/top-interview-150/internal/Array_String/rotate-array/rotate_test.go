package rotatearray

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestRotate(t *testing.T) {
	tests := []struct{
		nums []int
		k int
		want []int
	}{
		{[]int{1,2,3,4,5,6,7},3, []int{5,6,7,1,2,3,4}},
		{[]int{-1,-100,3,99},2, []int{3,99,-1,-100}},
		{[]int{-1},2,[]int{-1}},
	}

	for _, ts := range tests {
		testRotate := fmt.Sprintf("%v %v %v", ts.nums, ts.k, ts.want)
		t.Run(testRotate, func(t *testing.T) {
			got := rtt(ts.nums, ts.k)
			if !slices.Equal(got, ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

