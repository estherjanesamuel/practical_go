package lengthoflastword

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct{
		s string
		output int
		want string
	}{
		{"Hello World", 5 ,"World"},
		{"   fly me   to   the moon  ", 4 ,"moon"},
		{"luffy is still joyboy", 6 ,"joyboy"},
		{"The    quick    brown fox    jumps    over the    lazy dog", 3, "dog"},
	}

	for _, ts := range tests {
		testLengthOfLastWord := fmt.Sprintf("%v %v %v", ts.s, ts.output, ts.want)
		t.Run(testLengthOfLastWord, func(t *testing.T) {
			outputLength, lastWord := lengthOfLastWord(ts.s)
			if outputLength != ts.output  || lastWord != ts.want {
				t.Errorf("got %v & %v, wanted %v & %v", outputLength,lastWord, ts.output,ts.want)
			}
		})
	}
}
