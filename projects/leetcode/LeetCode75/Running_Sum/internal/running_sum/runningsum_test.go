package runningsum

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestRunningSum(t *testing.T) {
    tests := []struct{
        nums []int
        want []int
    }{
        {[]int{1,2,3,4},[]int{1,3,6,10}},
        {[]int{1,1,1,1,1},[]int{1,2,3,4,5}},
        {[]int{3,1,2,10,1},[]int{3,4,6,16,17}},
    }
    for _, tt := range tests {
        testRunningSum := fmt.Sprintf("%v", tt.nums)
        t.Run(testRunningSum, func(t *testing.T) {
            got := runningSum(tt.nums)
            // assert.Equal(t, tt.want, got)
            ok := slices.Equal(got,tt.want)
			if !ok {
				t.Errorf("got %v, wanted %v",got,tt.want)
			}
        })
    }
}


/*
Input: nums = []
Output: []
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
Example 2:

Input: nums = []
Output: []
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].
Example 3:

Input: nums = [3,1,2,10,1]
Output: []

*/