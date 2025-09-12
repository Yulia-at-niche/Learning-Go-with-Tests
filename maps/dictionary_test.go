package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("", func(t *testing.T) {
		dictionary := map[string]string{"test": "this is a test"}

		got := Search(dictionary, "test")
		want := "this is a test"

		if got != want {
			t.Errorf("got %q, wnat %q, given %q", got, want, "test")
		}

	})
}
