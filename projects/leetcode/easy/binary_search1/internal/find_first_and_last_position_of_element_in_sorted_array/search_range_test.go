package findfirstandlastpositionofelementinsortedarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert" // Add testify https://stackoverflow.com/questions/15311969/checking-the-equality-of-two-slices
)


func TestSearchRange(t *testing.T) {
	tests := []struct{
		nums []int
		target int
		want []int
	}{
		{[]int{5,7,7,8,8,10},8,[]int{3,4}},
		{[]int{5,7,7,8,8,10},6,[]int{-1,-1}},
		{[]int{},0,[]int{-1,-1}},
	}

	for _, ts := range tests {
		testSearchRange := fmt.Sprintf("%v, %v", ts.nums,ts.target)
		t.Run(testSearchRange, func(t *testing.T) {
			got := searchRange(ts.nums, ts.target)
			assert.Equal(t, ts.want, got) 
		})
		testSearchRangeTLE := fmt.Sprintf("%v, %v", ts.nums,ts.target)
		t.Run(testSearchRangeTLE, func(t *testing.T) {
			got := searchRange(ts.nums, ts.target)
			assert.Equal(t, ts.want, got) 
		})
	}
	
}