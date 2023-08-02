package reverselinkedlist

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestReverseList(t *testing.T) {
	tests := []struct{
		head,want []int
		
	}{
		{[]int{1,2,3,4,5},[]int{5,4,3,2,1}},
		{[]int{1,2},[]int{2,1}},
		{[]int{},[]int{}},
	}

	for _, ts := range tests {
		testReverseList := fmt.Sprintf("%v %v", ts.head, ts.want)
		t.Run(testReverseList, func(t *testing.T) {
			got := reverseList(ts.head)
			ok := slices.Equal(got,ts.want)
			if !ok {
				t.Errorf("got %v, wanted %v",got,ts.want)
			}
		})
	}
}
