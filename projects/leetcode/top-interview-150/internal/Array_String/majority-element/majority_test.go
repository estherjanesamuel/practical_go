package majorityelement

import (
	"fmt"
	"testing"
)

func TestMajority(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{3,2,3}, 3},
		{[]int{2,2,1,1,1,2,2}, 2},
		{[]int{2,1,1,1,2,1,2}, 1},
		{[]int{2,1,2,1,2,1,2}, 2},
		{[]int{1,2,3,1,2,3,1,2,3,1,2,3,1}, 1},
		{[]int{1,1,1,2,1,2,3,1,3,1,1,2,2,1,3,1}, 1},
	}
	for _, ts := range tests {
		testMajority := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testMajority, func(t *testing.T) {
			got := MajorityElement(ts.nums)
		if got != ts.want {
			t.Errorf("got %v, wanted %v",got, ts.want)
		}
		})
	}
}
