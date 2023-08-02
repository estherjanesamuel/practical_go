package math

import (
	"fmt"
	"testing"
)

func TestSqrtx(t *testing.T)  {
	tests := []struct{
		x, want int
	}{
		{4,2},
		{8,2},
		{9,3},
		{12,3},
		{16,4},
	}

	for _, ts := range tests {
		testSqrtx := fmt.Sprintf("%v %v", ts.x, ts.want)
		t.Run(testSqrtx, func(t *testing.T) {
			got := squareRoot(ts.x)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

/*
func TestSqrt(t *testing.T)  {
	tests := []struct{
		x, want int
	}{
		{4,16},
		{8,64},
	}

	for _, ts := range tests {
		testSqrt := fmt.Sprintf("%v %v", ts.x, ts.want)
		t.Run(testSqrt, func(t *testing.T) {
			got := sqrt(ts.x)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}


func TestPwr(t *testing.T)  {
	tests := []struct{
		x int 
		y float64
		want int
	}{
		{4,2,16},
		{8,2,64},
		{9,0.5,3},
	}

	for _, ts := range tests {
		testPwr := fmt.Sprintf("%v %v", ts.x, ts.want)
		t.Run(testPwr, func(t *testing.T) {
			got := pwr(ts.x, ts.y)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

*/