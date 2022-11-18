package kthmissingpositivenumber

import (
	"fmt"
	"testing"
)

func TestFindKthPositive(t *testing.T) {
	tests := []struct{
		arr []int
		k, want int
	}{
		{[]int{2,3,4,7,11},5,9},
		{[]int{1,2,3,4},2,6},
	}

	for _, ts := range tests {
		testFindKthPositive := fmt.Sprintf("%v %v %v", ts.arr, ts.k, ts.want)
		t.Run(testFindKthPositive, func(t *testing.T) {
			got := findKthPositive(ts.arr, ts.k)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
