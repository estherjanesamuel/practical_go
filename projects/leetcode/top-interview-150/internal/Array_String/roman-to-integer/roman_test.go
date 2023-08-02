package romantointeger

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct{
		s string
		want int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, ts := range tests {
		testRoman := fmt.Sprintf("%v %v", ts.s, ts.want)
		t.Run(testRoman, func(t *testing.T) {
			got := romanToInt(ts.s)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

func TestIntToRoman(t *testing.T) {
	tests := []struct{
		num int
		want string
	}{
		{3, "III"},
		{58,"LVIII", },
		{1994,"MCMXCIV", },
	}

	for _, ts := range tests {
		testRoman := fmt.Sprintf("%v %v", ts.num, ts.want)
		t.Run(testRoman, func(t *testing.T) {
			got := intToRmn(ts.num)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

