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

func TestAdd(t *testing.T) {
	t.Run("add a new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})
}
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should have found word:", err)
	}

	assertStringMatch(t, dictionary, got, definition)
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
