package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(3, 3)
	expected := 6

	if sum != expected {
		t.Errorf("expected %d got %d", expected, sum)
	}
}
