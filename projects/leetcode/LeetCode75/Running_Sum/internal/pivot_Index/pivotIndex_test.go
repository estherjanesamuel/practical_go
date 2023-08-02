package pivotindex

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{1,7,3,6,5,6}, 3},
		{[]int{1,2,3}, -1},
		{[]int{2,1,-1}, 0},
		{[]int{3,5,1,2,2,6,3,2},4},
		{[]int{1,2,3,4,5,2,7,8,9},6},
		{[]int{9,8,-2,6,5,4,3,2,1},3},
	}

	for _, ts := range tests {
		testPivotIndex := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testPivotIndex, func(t *testing.T) {
			got := pivotIndexII(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	} 
}
