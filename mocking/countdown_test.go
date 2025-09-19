package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

// We're making a mock of the real Sleep() method. Instead of it being a way to delay the program from running, it will increment the Call field of the SpySleeper struct
func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	// Countdown is going to call spySleeper.Sleep() instead of that delaying it's going to increment the Call field of the test's spySleeper
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper, want 3 got %d", spySleeper.Calls)
	}
}
