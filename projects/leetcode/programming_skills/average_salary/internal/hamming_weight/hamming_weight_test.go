package hammingweight

import (
	"fmt"
	"testing"
)


func TestHammingWeight(t *testing.T) {
	tests := []struct{
		input uint32
		want int
	}{
		{00000000000000000000000000001011, 3},
		{00000000000000000000000010000000, 1},
		{4294967293, 31},
	}
	for _, ts := range tests {
		testHammingWeeight := fmt.Sprintf("%v %v",ts.input, ts.want)
		t.Run(testHammingWeeight, func(t *testing.T) {
			got := hammingWeight(ts.input)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}

func TestHammingWeight2(t *testing.T) {
	tests := []struct{
		input uint32
		want uint32
	}{
		{11, 3},
		{128, 1},
		{4294967293, 31},
	}
	for _, ts := range tests {
		testHammingWeeight2 := fmt.Sprintf("%v %v",ts.input, ts.want)
		t.Run(testHammingWeeight2, func(t *testing.T) {
			got := hammingWeight2(uint32(ts.input))
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}