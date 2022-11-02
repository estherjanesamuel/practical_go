package hello

import "testing"

func TestHello(t *testing.T) {
	// Here we are introducing another tool in our testing arsenal, subtests. 
	// Sometimes it is useful to group tests around a "thing" and then 
	// have subtests describing different scenarios.
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("arief")
		want := "hello, arief"
		assertCorrectMessage(t,got, want)
	})
	t.Run("say 'hello, world' when an empty string is supplied",func(t *testing.T) {
		got := Hello("")
		want := "hello, world"
		assertCorrectMessage(t,got, want)
	}) 
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}