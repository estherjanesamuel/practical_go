package findnearestpointthathasthesamexycoordinate

import "testing"

func TestNearestValidPoint(t *testing.T) {
	tests := []struct{
		x,y int
		points [][]int
		want int
	}{
		{3,4,[][]int{[]int{1,3,2,2,4},[]int{2,1,4,3,4}},2},
		{3,4,[][]int{[]int{3},[]int{4}},0},
		{3,4,[][]int{[]int{2},[]int{3}},-1},
	}
}