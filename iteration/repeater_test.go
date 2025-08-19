package iteration

import "testing"

func TestRepeater(t *testing.T) {
	repeating := Repeater("x")
	expected := "xxxxx"

	if repeating != expected {
		t.Errorf("Expected %q, got %q", expected, repeating)
	}
}
