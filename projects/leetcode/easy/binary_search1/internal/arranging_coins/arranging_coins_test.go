package arrangingcoins

import (
	"fmt"
	"testing"
)

func TestArrangingCoins(t *testing.T) {
	tests := []struct{
		n,  want int
	}{
		{5,2},
		{8,3},
		{9,3},
		{1000,44},
	}
	for _, ts := range tests {
		testArrangingCoins := fmt.Sprintf("%v", ts.n)
		t.Run(testArrangingCoins, func(t *testing.T) {
			got := arrangeCoins(ts.n)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
		// testArrCoins := fmt.Sprintf("%v %v", ts.n, ts.want)
		// t.Run(testArrCoins, func(t *testing.T) {
		// 	got := arrCoins(ts.n)
		// 	if got != ts.want {
		// 		t.Errorf("got %v, wanted %v", got, ts.want)
		// 	}
		// })
	}
}
