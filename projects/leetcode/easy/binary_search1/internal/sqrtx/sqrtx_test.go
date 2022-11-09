package sqrtx

import (
	"fmt"
	"testing"
)

func TestSqrt(t *testing.T) {
	tests := []struct{
		x, want int
	}{
		{4,2},
		{8,2},
	}

	for _, ts := range tests {
		testmysqrt := fmt.Sprintf("mySqrt(%v)", ts.x)
		t.Run(testmysqrt, func(t *testing.T) {
			got := mySqrt(ts.x)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
		testsqrt := fmt.Sprintf("sqrt(%v)", ts.x)
		t.Run(testsqrt, func(t *testing.T) {
			got := sqrt(ts.x)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}