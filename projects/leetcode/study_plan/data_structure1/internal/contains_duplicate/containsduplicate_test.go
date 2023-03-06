package containsduplicate

import (
	"fmt"
	"testing"
)

/*
	given an array nums, return true
	if any value appears at least twice
	in the array, and return false if
	every element is distinct
*/
func TestContainsDuplicate(t *testing.T) {
	tests := []struct{
		nums []int
		want bool
	}{
		{[]int{1,2,3,1},true},
		{[]int{1,2,3,4},false},
		{[]int{1,2,3,4,2},true},
		{[]int{1,2,3,4,5,6,7,8,9,10},true},

	}

	for _, ts := range tests {
		testContainsDuplicate := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testContainsDuplicate, func(t *testing.T) {
			got := containsDuplicate(ts.nums)

			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}


func BenchmarkContainsDuplicate(b *testing.B) {
	table := []struct{
		nums []int
		want bool
	}{
		{[]int{1,2,3,1},true},
		{[]int{1,2,3,4},false},
		{[]int{1,2,3,4,2},true},
		{[]int{1,2,3,4,5,6,7,8,9,10},true},

	}
	for _, v := range table{
		b.Run(fmt.Sprintf("containsDuplicate_%v", v.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				containsDuplicate(v.nums)
			}
		})

		b.Run(fmt.Sprintf("containsDuplicateII_%v", v.nums), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				containsDuplicateII(v.nums)
			}
		})
		
	}
	
}