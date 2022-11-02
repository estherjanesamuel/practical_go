package hello

import "testing"

func TestXxx(t *testing.T) {
	got := Hello("arief")
	want := "hello, arief"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}