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
		_, err := dictionary.Search("unknown")
		want := "could not find what you are looking for"

		if err == nil {
			t.Fatal("expected error, got none")
		}

		assertStringMatch(t, dictionary, err.Error(), want)

	})
}

func assertStringMatch(t testing.TB, given map[string]string, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, given)
	}
}
