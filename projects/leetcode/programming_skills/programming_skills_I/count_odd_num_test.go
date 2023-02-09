package programmingskillsi

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCountOdds(t *testing.T) {
	tests := []struct{
		low,high int
		want []int
	}{
		{3,7,[]int{3,5,7}},
		{8,10,[]int{9}},
	}

	for _, ts := range tests {
		testCountOdds := fmt.Sprintf("%v, %v, %v", ts.low, ts.high, ts.want)
		t.Run(testCountOdds, func(t *testing.T) {
			got := countOdds(ts.low, ts.high)
			if !cmp.Equal(got, ts.want){
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

func TestCountOdd(t *testing.T) {
	tests := []struct{
		low,high,want int
	}{
		{3,7,3},
		{8,10,1},
	}

	for _, ts := range tests {
		testCountOdds := fmt.Sprintf("%v, %v, %v", ts.low, ts.high, ts.want)
		t.Run(testCountOdds, func(t *testing.T) {
			got := countOdd(ts.low, ts.high)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}

			result := countOdd1(ts.low, ts.high)
			if result != ts.want{
				t.Errorf("got %v, wanted %v", got, ts.want)
			}

			result2 := countOdd2(ts.low, ts.high)
			if result2 != ts.want{
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
		
	}
}

func TestCountOdds3(t *testing.T) {
	tests := []struct{
		low,high,want int
	}{
		{3,7,3},
		{8,10,1},
	}

	for _, ts := range tests {
		testCountOdds3 := fmt.Sprintf("%v, %v, %v", ts.low, ts.high, ts.want)
		t.Run(testCountOdds3, func(t *testing.T) {
			got := countOdds3(ts.low, ts.high)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
}}