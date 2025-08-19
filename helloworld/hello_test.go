package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Narla", "")
		want := "Hello, Narla"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Juanes", "Spanish")
		want := "Hola, Juanes"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello in Klingon", func(t *testing.T) {
		got := Hello("Kahless", "Klingon")
		want := "nuqneH, Kahless"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper() // This tells the test suite the this method is a helper. Doing this will make test failures report the line number it failed on in TestHello() (EX: line 14) instead of here in assertCorrectMessage() (EX: line 19)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
