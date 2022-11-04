package peakindexinamountainarray

import (
	"fmt"
	"testing"
)

func TestPeakIndexInAMountainArray(t *testing.T) {
	t.Run("first solution", func(t *testing.T) {
		tests := []struct {
			arr  []int
			want int
		}{
			{[]int{0, 1, 0}, 1},
			{[]int{0, 2, 1, 0}, 1},
			{[]int{0, 10, 5, 2}, 1},
			{[]int{3,4,5,1}, 2},
		}

		for _, tt := range tests {
			testpeakindex := fmt.Sprintf("%v", tt.arr)
			t.Run(testpeakindex, func(t *testing.T) {
				got := peakIndexInMountainArray_1(tt.arr)
				if got != tt.want {
					t.Errorf("got %d, wanted %d", got, tt.want)
				}
			})
		}
	})

	t.Run("second solution", func(t *testing.T) {
		tests := []struct {
			arr  []int
			want int
		}{
			{[]int{0, 1, 0}, 1},
			{[]int{0, 2, 1, 0}, 1},
			{[]int{0, 10, 5, 2}, 1},
			{[]int{3,4,5,1}, 2},
		}

		for _, tt := range tests {
			testpeakindex := fmt.Sprintf("%v", tt.arr)
			t.Run(testpeakindex, func(t *testing.T) {
				got := peakIndexInMountainArray(tt.arr)
				if got != tt.want {
					t.Errorf("got %d, wanted %d", got, tt.want)
				}
			})
		}
	})
}
