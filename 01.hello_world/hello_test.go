package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Fisk")
	want := "Hello, Fisk"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}