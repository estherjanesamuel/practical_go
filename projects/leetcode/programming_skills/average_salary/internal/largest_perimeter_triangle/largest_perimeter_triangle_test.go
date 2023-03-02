package largestperimetertriangle

import (
	"fmt"
	"testing"
)

func TestLrgestPerimeterTriangle(t *testing.T) {
	tests := []struct{
		nums []int
		want int
	}{
		{[]int{2,1,2},5},
		{[]int{1,2,1,10},0},
	}

	for _, ts := range tests {
		testTestLrgestPerimeterTriangle := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testTestLrgestPerimeterTriangle, func(t *testing.T) {
			got := largestPerimeter(ts.nums)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}