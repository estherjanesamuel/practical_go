package maximumsubarray

import (
	"fmt"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{-2,1,-3,4,-1,2,1,-5,4},6},
		{[]int{1},1},
		{[]int{-1},-1},
		{[]int{-2,-1},-1},
		{[]int{5,4,-1,7,8},23}, //[-2,-1]
	}

	for _, ts := range tests {
		testMaxSubarray := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testMaxSubarray, func(t *testing.T) {
			got := maximumSubarray(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v",got, ts.want)
			}
		})
	} 
}


func TestMaximumSubarrayII(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{-2,1,-3,4,-1,2,1,-5,4},6},
		{[]int{1},1},
		{[]int{-1},-1},
		{[]int{-2,-1},-1},
		{[]int{5,4,-1,7,8},23},
	}

	for _, ts := range tests {
		testMaxSubarrayII := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testMaxSubarrayII, func(t *testing.T) {
			got := maximumSubarrayII(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v",got, ts.want)
			}
		})
	} 
}
/*
func TestMin(t *testing.T) {
	tests := []struct{
		x,y int
		want int
	}{
		{10,9,9},
		{9,10,9},
		{1,2,1},
	}
	for _, ts := range tests {
		testMin := fmt.Sprintf("%v %v %v", ts.x, ts.y,ts.want)
		t.Run( testMin , func(t *testing.T) {
			got := min(ts.x,ts.y)
			if got != ts.want {
				t.Errorf("got %v, wanted %v",got, ts.want)
			}
		})
	} 
}


func TestMax(t *testing.T) {
	tests := []struct{
		x,y int
		want int
	}{
		{10,9,10},
		{9,10,10},
		{1,2,2},
	}
	for _, ts := range tests {
		testMax := fmt.Sprintf("%v %v %v", ts.x, ts.y,ts.want)
		t.Run( testMax , func(t *testing.T) {
			got := max(ts.x,ts.y)
			if got != ts.want {
				t.Errorf("got %v, wanted %v",got, ts.want)
			}
		})
	} 
}
*/