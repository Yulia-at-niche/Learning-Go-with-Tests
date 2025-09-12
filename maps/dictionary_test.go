package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is a test"}

		got := Search(dictionary, "test")
		want := "this is a test"

		assertStringMatch(t, dictionary, want, got)

	})
}

func assertStringMatch(t testing.TB, given map[string]string, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, wnat %q, given %q", got, want, "test")
	}
}
