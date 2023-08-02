package Hashmap

import (
	"fmt"
	"testing"
)

func TestContainsDuplicate(t *testing.T)  {
	tests := []struct{
		nums []int
		k int
		want bool
	}{
		// {[]int{1,2,3,1}, 3, true},
		// {[]int{1,0,1,1}, 1, true},
		{[]int{1,2,3,1,2,3}, 2, false},
		// {[]int{1,2,1,1,2,3}, 2, true},
	}

	for _, ts := range tests {
		testContainsDuplicate := fmt.Sprintf("%v %v %v", ts.nums, ts.k, ts.want)
		t.Run(testContainsDuplicate, func(t *testing.T) {
			got := nearby_distinc(ts.nums, ts.k)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

