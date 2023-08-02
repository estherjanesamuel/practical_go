package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	tests := []struct{
		input, want int
	}{
		{20,6765},
		{19,4181},
		{10,55},
		{9,34},
		{7,13},
		{6,8},
		{5,5},
		{4,3},
		{3,2},
		{2,1},
	}

	for _, tt:= range tests {
		fibTest := fmt.Sprintf("%d %d", tt.input, tt.want)
		t.Run(fibTest, func(t *testing.T) {
			got := fib(tt.input)

			if got != tt.want {
				t.Errorf("got %v, wanted %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(20)
	}	
}
