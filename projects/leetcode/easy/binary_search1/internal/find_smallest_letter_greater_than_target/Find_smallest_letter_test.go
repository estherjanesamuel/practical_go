package findsmallestlettergreaterthantarget

import (
	"fmt"
	"testing"
)

func TestFindSmallestLetter(t *testing.T) {
	tests := []struct{
		letters []string
		target string
		want string
	}{
		{[]string{"c","f","j"},"a","c"},
		{[]string{"c","f","j"},"c","f"},
		{[]string{"x","x","y","y"},"z","x"},
		{[]string{"c","f","j"},"d","f"},
	}

	for _, ts := range tests {
		testFindSmallest := fmt.Sprintf("%v, %v", ts.letters, ts.target)
		t.Run(testFindSmallest, func(t *testing.T) {
			got := findSmallestLetter(ts.letters, ts.target)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
		
		
		// testNextGreaterLetter := fmt.Sprintf("%v, %v", ts.letters, ts.target)
		// t.Run(testNextGreaterLetter, func(t *testing.T) {
		// 	got := nextGreaterLetter(ts.letters, ts.target)
		// 		if got != ts.want {
		// 			t.Errorf("got %v, wanted %v", got, ts.want)
		// 		}
		// })
		// testNextGreatestLetter := fmt.Sprintf("%v, %v", ts.letters, ts.target)
		// t.Run(testNextGreatestLetter, func(t *testing.T) {
		// 	got := nextGreaterLetter(ts.letters, ts.target)
		// 		if got != ts.want {
		// 			t.Errorf("got %v, wanted %v", got, ts.want)
		// 		}
		// })
		
	}
}

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct{
		letters []byte
		target byte
		want byte
	}{
		{[]byte{'c','f','j'},'a','c'},
		{[]byte{'c','f','j'},'c','f'},
		{[]byte{'x','x','y','y'},'z','x'},
		{[]byte{'c','f','j'},'d','f'},
	}
	for _, ts := range tests {
		// testNextGreatestLetter := fmt.Sprintf("%v, %v", ts.letters, ts.target)
		// t.Run(testNextGreatestLetter, func(t *testing.T) {
		// 	got := nextGreatestLetter(ts.letters, ts.target)
		// 		if got != ts.want {
		// 			t.Errorf("got %v, wanted %v", got, ts.want)
		// 		}
		// })
		testNextGreaterLetterBytes := fmt.Sprintf("%v, %v", ts.letters, ts.target)
		// t.Run(testNextGreaterLetterBytes, func(t *testing.T) {
		// 	got := nextGreaterLetterBytes(ts.letters, ts.target)
		// 	if got != ts.want {
		// 		t.Errorf("got %v, wanted %v", got, ts.want)
		// 	}
		// })

		t.Run(testNextGreaterLetterBytes, func(t *testing.T) {
			got := findNextGreatesLetter(ts.letters, ts.target)
			if got != ts.want {
				t.Errorf("got %v, wanted %v", got, ts.want)
			}
		})
	}
}