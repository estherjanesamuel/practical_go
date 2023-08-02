package reversearray

import (
	"fmt"
	"testing"

	"golang.org/x/exp/slices"
)

func TestReverse(t *testing.T) {
	tests := []struct{
		nums []int
		want []int
	}{
		{[]int{1,2,3,4,5,6,7,8,9}, []int{9,8,7,6,5,4,3,2,1}},
	}

	for _, ts := range tests {
		testReverse := fmt.Sprintf("%v %v", ts.nums, ts.want)
		t.Run(testReverse, func(t *testing.T) {	
			got := reverse(ts.nums)
			if !slices.Equal(got,ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}


}

func TestReverseII(t *testing.T) {
	tests := []struct{
		nums []int
		start , end int
		want []int
	}{
		{[]int{1,2,3,4,5,6,7,8,9},0,8,[]int{9,8,7,6,5,4,3,2,1}},
		{[]int{1,2,3,4,5,6,7,8,9},1,8,[]int{1,9,8,7,6,5,4,3,2}},
		// {[]int{1,2,3,4,5,6,7,8,9},2,9,[]int{1,2,9,8,7,6,5,4,3}},
		// {[]int{1,2,3,4,5,6,7,8,9},3,9,[]int{1,2,3,9,8,7,6,5,4}},
		// {[]int{1,2,3,4,5,6,7,8,9},4,9,[]int{1,2,3,4,9,8,7,6,5}},
		// {[]int{1,2,3,4,5,6,7,8,9},5,9,[]int{1,2,3,4,5,9,8,7,6}},
		// {[]int{1,2,3,4,5,6,7,8,9},6,9,[]int{1,2,3,4,5,6,9,8,7}},
		// {[]int{1,2,3,4,5,6,7,8,9},7,9,[]int{1,2,3,4,5,6,7,9,8}},
		// {[]int{1,2,3,4,5,6,7,8,9},8,9,[]int{1,2,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},9,9,[]int{1,2,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,1,[]int{1,2,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,2,[]int{2,1,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,3,[]int{3,2,1,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,4,[]int{4,3,2,1,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,5,[]int{5,4,3,2,1,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,6,[]int{6,5,4,3,2,1,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,7,[]int{7,6,5,4,3,2,1,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,8,[]int{8,7,6,5,4,3,2,1,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,9,[]int{9,8,7,6,5,4,3,2,1}},
		// {[]int{1,2,3,4,5,6,7,8,9},3,6,[]int{9,8,7,6,5,4,3,2,1}},
	}

	for _, ts := range tests {
		testReverseII := fmt.Sprintf("%v %v %v %v", ts.nums,ts.start, ts.end, ts.want)
		t.Run(testReverseII, func(t *testing.T) {	
			got := reverseII(ts.nums,ts.start,ts.end)
			if !slices.Equal(got,ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}


}

func TestReverseIII(t *testing.T) {
	tests := []struct{
		nums []int
		start , end int
		want []int
	}{
		{[]int{1,2,3,4,5,6,7,8,9},8,7,[]int{9,1,2,3,4,5,6,7,8}},
		// {[]int{1,2,3,4,5,6,7,8,9},1,8,[]int{1,9,8,7,6,5,4,3,2}},
		// {[]int{1,2,3,4,5,6,7,8,9},2,9,[]int{1,2,9,8,7,6,5,4,3}},
		// {[]int{1,2,3,4,5,6,7,8,9},3,9,[]int{1,2,3,9,8,7,6,5,4}},
		// {[]int{1,2,3,4,5,6,7,8,9},4,9,[]int{1,2,3,4,9,8,7,6,5}},
		// {[]int{1,2,3,4,5,6,7,8,9},5,9,[]int{1,2,3,4,5,9,8,7,6}},
		// {[]int{1,2,3,4,5,6,7,8,9},6,9,[]int{1,2,3,4,5,6,9,8,7}},
		// {[]int{1,2,3,4,5,6,7,8,9},7,9,[]int{1,2,3,4,5,6,7,9,8}},
		// {[]int{1,2,3,4,5,6,7,8,9},8,9,[]int{1,2,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},9,9,[]int{1,2,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,1,[]int{1,2,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,2,[]int{2,1,3,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,3,[]int{3,2,1,4,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,4,[]int{4,3,2,1,5,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,5,[]int{5,4,3,2,1,6,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,6,[]int{6,5,4,3,2,1,7,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,7,[]int{7,6,5,4,3,2,1,8,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,8,[]int{8,7,6,5,4,3,2,1,9}},
		// {[]int{1,2,3,4,5,6,7,8,9},0,9,[]int{9,8,7,6,5,4,3,2,1}},
		// {[]int{1,2,3,4,5,6,7,8,9},3,6,[]int{9,8,7,6,5,4,3,2,1}},
	}

	for _, ts := range tests {
		testReverseIII := fmt.Sprintf("%v %v %v %v", ts.nums,ts.start, ts.end, ts.want)
		t.Run(testReverseIII, func(t *testing.T) {	
			got := reverseIII(ts.nums,ts.start,ts.end)
			if !slices.Equal(got,ts.want) {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}


}


func TestRvrs(t *testing.T) {
	tests := []struct{
		str string
		want string
	}{
		{"raceacar","racaecar"},
		{"abcdefg", "gfedcba"},
	}

	for _, ts := range tests {
		testReverse := fmt.Sprintf("%v %v", ts.str, ts.want)
		t.Run(testReverse, func(t *testing.T) {	
			// got := rvrs(ts.str, 0, len(ts.str)-1)
			got := rvrsII(ts.str)
			fmt.Println(ts.str,got,ts.want)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}


}