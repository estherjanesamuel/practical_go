package happynumber

import (
	"fmt"
	"testing"
)

func TestHappyNumber(t *testing.T)  {
	tests := []struct{
		n int
		want bool
	}{
		{19,true},
		// {2, false},
	}

	 for _, ts := range tests {
		testHappyNumber := fmt.Sprintf("%v %v", ts.n, ts.want)
		t.Run(testHappyNumber, func(t *testing.T) {
			got := is_happy(ts.n)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	 }
}

