package searchinsert

import (
	"fmt"
	"testing"
)

func TestSearchInsert(t *testing.T) {
	t.Run("first attempt to searchInsert problem on leetcode", func(t *testing.T) {
		tests := []struct {
			nums         []int
			target, want int
		}{
			{[]int{1, 3, 5, 6}, 5, 2},
			{[]int{1, 3, 5, 6}, 2, 1},
			{[]int{1, 3, 5, 6}, 7, 4},
			{[]int{1, 3, 5, 6}, 0, 0},
		}

		for _, tt := range tests {
			testsearchinsert := fmt.Sprintf("%v, %v", tt.nums, tt.target)
			t.Run(testsearchinsert, func(t *testing.T) {
				got := searchInsert(tt.nums, tt.target)
				assertCorrectResult(t,got,tt.want)
			})
		}
	})
	t.Run("second attempt to searchInsert problem on leetcode", func(t *testing.T) {
		tests := []struct {
			nums         []int
			target, want int
		}{
			{[]int{1, 3, 5, 6}, 5, 2},
			{[]int{1, 3, 5, 6}, 2, 1},
			{[]int{1, 3, 5, 6}, 7, 4},
			{[]int{1, 3, 5, 6}, 0, 0},
		}

		for _, tt := range tests {
			testsearchinsert := fmt.Sprintf("%v, %v", tt.nums, tt.target)
			t.Run(testsearchinsert, func(t *testing.T) {
				got := searchinsert(tt.nums, tt.target)
				assertCorrectResult(t,got,tt.want)
			})
		}
	})
}

func assertCorrectResult(t testing.TB, got,want int)  {
	t.Helper()
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}