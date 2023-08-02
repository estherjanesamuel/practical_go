package removeduplicatesfromsortedarray

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{1,1,2},2},
		{[]int{0,0,1,1,1,2,2,3,3,4},5},
	}

	for _, ts := range tests {
		testRemove := fmt.Sprintf("%v %v", ts.nums,ts.want)
		t.Run(testRemove, func(t *testing.T) {
			got := Remove(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}


func TestRemoveDuplicate(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{1,1,1,2,2,2,3},5},
		// {[]int{0,0,1,1,1,1,2,3,3},7},
	}

	for _, ts := range tests {
		testRemove := fmt.Sprintf("%v %v", ts.nums,ts.want)
		t.Run(testRemove, func(t *testing.T) {
			got := removeDuplicate(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

func TestRemove_Duplicate(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{1,1,1,2,2,2,3},5},
		{[]int{0,0,1,1,1,1,2,3,3},7},
	}

	for _, ts := range tests {
		testRemove := fmt.Sprintf("%v %v", ts.nums,ts.want)
		t.Run(testRemove, func(t *testing.T) {
			got := remove_duplicate(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

