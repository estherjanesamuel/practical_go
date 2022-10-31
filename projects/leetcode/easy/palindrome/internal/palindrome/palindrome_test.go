package palindrome

import (
	"fmt"
	"testing"
)


func TestPalindrome(t *testing.T) {
	got := Palindrome(-121)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestPalindromeTableDriven(t *testing.T) {
var tests = []struct {
	x int
	want bool
}{
	{121,true},
	{-121,false},
	{10,false},
	{1001,true},
}
for _, tt := range tests {
	testname := fmt.Sprintf("%d", tt.x)
	t.Run(testname, func (t *testing.T)  {
		ans := Palindrome(tt.x)
		if ans != tt.want {
			t.Errorf("got %v, wanted %v", ans, tt.want)
		}
	})
}
}

func TestReverse(t *testing.T) {
	got := Reverse("arief")
	want := "feira"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestReverseStr(t *testing.T) {
	got := ReverseStr("1001")
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	got := Palindrome(121)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func BenchmarkReverseStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("str")
	}	
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseStr("1001")
	}	
}