package structs

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter", func(t *testing.T) {
		got := Perimeter(10.0, 10.0)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	got := Area(10.0, 5.5)
	want := 55.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
