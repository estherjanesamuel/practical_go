package intervals

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		// {[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
		// {[]int{-2147483648,-2147483647,2147483647},[]string{"0","2->4","6","8->9"}},
	}

	for _, ts := range tests {
		testSummaryRanges := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testSummaryRanges, func(t *testing.T) {
			got := summary_ranges(ts.nums)
			if !slices.Equal(got, ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
