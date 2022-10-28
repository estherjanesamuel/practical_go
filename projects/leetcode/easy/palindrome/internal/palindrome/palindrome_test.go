package palindrome

import "testing"


func TestPalindrome(t *testing.T) {
	got := Palindrome(-121)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestReverse(t *testing.T) {
	got := Reverse("arief")
	want := "feira"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestIsPalindrome(t *testing.T) {
	got := Palindrome(121)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}