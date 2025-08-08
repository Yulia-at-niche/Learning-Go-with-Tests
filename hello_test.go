package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Narla")
	want := "Hello, Narla"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
