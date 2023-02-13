package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Fisk")
		want := "Hello, Fisk"
		assertCorrectMessage(t,got,want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T){
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t,got,want)
	})
}

// Helper function to reduce repeated code in subtests above
func assertCorrectMessage(t testing.TB, got, want string) {
	// t.Helper we point error at parent test
	t.Helper()
	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}