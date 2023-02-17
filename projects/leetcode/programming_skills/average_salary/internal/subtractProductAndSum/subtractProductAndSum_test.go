package subtractproductandsum

import (
	"fmt"
	"testing"
)

func TestSubtractProductAndSum(t *testing.T) {
	tests := []struct{
		n, want int
	}{
		{234,15},
		{4421, 21},
	}
	for _, ts := range tests {
		testSubtractProductAndSum := fmt.Sprintf("%v %v", ts.n, ts.want)
		t.Run(testSubtractProductAndSum, func(t *testing.T) {
			got := subtractProductAndSum(ts.n)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

