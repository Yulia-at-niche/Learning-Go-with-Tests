package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStringMatch(t, dictionary, want, got)

	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		want := ErrWordNotFound

		if got == nil {
			t.Fatal("expected error, got none")
		}

		assertError(t, got, want)
	})
}

func assertStringMatch(t testing.TB, given map[string]string, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, given)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected error, did not get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
