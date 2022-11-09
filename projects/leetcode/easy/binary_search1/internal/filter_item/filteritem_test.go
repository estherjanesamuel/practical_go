package filteritem

import "testing"

func TestFilterItem(t *testing.T) {
	arr := [][]int{{5,2},{4,2},{7,7},{8,3},{4,5}}
	query := [][]int{{5,5,2,2,},{4,7,5,7},{4,8,2,7}}
	i,j,k := FilterItem(5, arr, 3, query)
	wanti := 1
	wantj := 2
	wantk := 3
	if i != wanti && j != wantj && k != wantk {
		t.Errorf("got %v,%v,%v , wanted %v,%v,%v", i,j,k, wanti,wantj,wantk)
	}
}
