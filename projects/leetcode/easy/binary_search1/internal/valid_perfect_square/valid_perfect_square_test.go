package validperfectsquare

import (
	"fmt"
	"testing"
)

func TestValidPerfectSquare(t *testing.T) {
	tests := []struct{
		num int
		want bool
	}{
		{16,true},
		{12,false},
		{10,false},
		{104976,true},
		{1,true},
	}
	for _, ts := range tests {
		testname := fmt.Sprintf("%v", ts.num)
		t.Run(testname, func(t *testing.T) {
			got := isPerfectSquare2(ts.num)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}

}

// go test -bench=. -count 5 -run=^#
func BenchmarkValidPerfectSquare(b *testing.B) {
	b.Run("benchmark isPerfectSquare2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPerfectSquare2(104976)
		}
	}) 
	b.Run("benchmark PerfectSquare", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			perfectSquare(1049761231249124132)
		}
	}) 
	b.Run("benchmark isPerfectSquare", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPerfectSquare(1049761231249124132)
		}
	}) 
	b.Run("benchmark isPerfectSquare2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPerfectSquare2(1049761231249124132)
		}
	})
	b.Run("benchmark isPerfectSquare3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPerfectSquare3(1049761231249124132)
		}
	}) 
	b.Run("benchmark isPerfectSquare4", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isPerfectSquare4(1049761231249124132)
		}
	}) 
}