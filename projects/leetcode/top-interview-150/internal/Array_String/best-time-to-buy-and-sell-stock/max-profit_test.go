package array_string

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	} {
		{[]int{2,1,2,1,0,1,2}, 2},
		{[]int{2,1,2,1,0,1,2}, 2},
		{[]int{7,1,5,3,6,4}, 5},
		{[]int{3,2,6,5,0,3}, 4},
		{[]int{2,4,1}, 2},
		{[]int{7,6,4,3,1}, 0},

		
	}

	for _, ts := range tests {
		testMaxProfit := fmt.Sprintf("%v, %v", ts.nums, ts.want)
		t.Run(testMaxProfit, func(t *testing.T) {
			got := maxCuan(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	} {
		{[]int{7,1,5,3,6,4}, 7},
		{[]int{1,2,3,4,5}, 4},
		{[]int{7,6,4,3,1}, 0},
		{[]int{2,1,2,1,0,1,2}, 3},
		

		
	}

	for _, ts := range tests {
		testMaxProfit := fmt.Sprintf("%v, %v", ts.nums, ts.want)
		t.Run(testMaxProfit, func(t *testing.T) {
			got := maxCuanII(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
