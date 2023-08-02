package removeelement

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct{
		nums []int
		val, want int
	}{
		{[]int{3,2,2,3},3,2},
		{[]int{0,1,2,2,3,0,4,2},2,5},
	}

	for _, ts := range tests {
		testRemove := fmt.Sprintf("%v %v %v", ts.nums,ts.val,ts.want)
		t.Run(testRemove, func(t *testing.T) {
			got := Remove(ts.nums,  ts.val)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

