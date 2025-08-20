package iteration

import "testing"

func TestRepeater(t *testing.T) {
	repeating := Repeater("x", 3)
	expected := "xxx"

	if repeating != expected {
		t.Errorf("Expected %q, got %q", expected, repeating)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeater("a", 2)
	}
}
