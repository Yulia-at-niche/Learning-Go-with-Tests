package dependancyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Narla")

	got := buffer.String()
	want := "Hello, Narla"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
