package longestcommonprefix

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct{
		strs []string
		want string
	}{
		{[]string{"flower","flow","flight"}, "fl"},
		{[]string{"dog","racecar","car"}, ""},
		{[]string{"mint", "mini", "mineral"}, "min"},
	}

	for _, ts := range tests {
		testLongestCommonPrefix := fmt.Sprintf("%v %v", ts.strs, ts.want)
		t.Run(testLongestCommonPrefix, func(t *testing.T) {
			got := longestCommonPre(ts.strs)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}
