package guess

import (
	"fmt"
	"testing"
)

func TestGuess(t *testing.T) {
	got := GuessNumber(2)
	want := 1

	if got != want {
		t.Errorf("got %v, wanted %v",got, want)
	}
}

func TestGuessTableDriven(t *testing.T) {
	var tests = []struct{
		n,pick int
		want int
	}{
		{10,6,6},
		{2,1,1},
		{1,1,1},
		{100,31,31},
		{1000,31,31},
		{10000,31,31},
		{100000,31,31},
		{1000000,31,31},
		{10000000,31,31},
		{100000000,31,31},
		{1000000000,31,31},
		{10000000000,31,31},
		{100000000000,31,31},
	}

	for _, tt := range tests {
		guestTest := fmt.Sprintf("%d, %d", tt.n, tt.pick)
		t.Run(guestTest, func (t *testing.T)  {
			ans := guessNumber(tt.n, tt.pick)
			if ans != tt.want {
				t.Errorf("got %v, want %v", ans, tt.want)
			}
		})
	}
}

//go test -bench=. | go test -bench=. -count 5 -run=^#
func BenchmarkMid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mid(100000000000)
	}
}

func BenchmarkMid_Without(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mid(100000000000)
	}
}