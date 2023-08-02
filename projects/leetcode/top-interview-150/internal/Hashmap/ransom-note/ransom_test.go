package Hashmap

import (
	"fmt"
	"testing"
)

func TestRansom(t *testing.T) {
	tests := []struct{
		ransomNote string
		magazine string
		want bool
	}{
		{"a", "b", false},
		{"aa", "aab", true},
		{"aaab", "aab", false},
		{"aab","baa", true},
		{"zzxy","yxzz", true},
		{"aa", "ab", false},
		{"bg","efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj",true},
	}

	for _, ts := range tests {
		testRansom := fmt.Sprintf("%v %v %v", ts.ransomNote, ts.magazine, ts.want)
		t.Run(testRansom, func(t *testing.T) {
			got := construct(ts.ransomNote, ts.magazine)
			if got != ts.want {
				t.Errorf("got %v, wanted %v",got, ts.want)
			}
		})
	}
}
