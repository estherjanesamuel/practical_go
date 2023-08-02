package twosum

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestTwoSum(t *testing.T)  {
	tests := []struct{
		nums []int
		target int
		want []int
	}{
		{[]int{3,3}, 6,[]int{0,1}},
		{[]int{3,2,4}, 6,[]int{1,2}},
		{[]int{2,7,11,15}, 9,[]int{0,1}},
		{[]int{3,2,3}, 6,[]int{0,2}},
	}
	for _, ts := range tests {
		testTwoSum := fmt.Sprintf("%v, %v, %v", ts.nums, ts.target, ts.want)
		t.Run(testTwoSum, func(t *testing.T) {
			got := sumTwo(ts.nums, ts.target)
			if !slices.Equal(got, ts.want) {
				t.Errorf("got, %v. wanted %v",got, ts.want)
			}
		})
	}
}

func BenchmarkTwoSum(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		twoSum([]int{3,2,4}, 6)
	}
}

func BenchmarkSumTwo(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		sumTwo([]int{3,2,4}, 6)
	}
}