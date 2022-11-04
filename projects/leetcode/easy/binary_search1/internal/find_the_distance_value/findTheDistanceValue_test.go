package findthedistancevalue

import "testing"

// func TestFindDistanceValueSingle(t *testing.T) {
// 	got := findTheDistanceValue([]int{4, 5, 8}, []int{10, 9, 1, 8}, 2)
// 	want := 2
// 	if got != want {
// 		t.Errorf("got %v, wanted %v", got, want)
// 	}
// }

func TestFindDistanceValue(t *testing.T) {
	tests := []struct {
		num, arr       []int
		distance, want int
	}{
		{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2, 2},
		{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3, 2},
		{[]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, -7}, 6, 1},
	}

	for _, ts := range tests {
		got := findTheDistanceValue(ts.num, ts.arr, ts.distance)
		if got != ts.want {
			t.Errorf("got %v, wanted %v", got, ts.want)
		}
	}
}
