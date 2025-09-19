package main

import (
	"bytes"
	"reflect"
	"testing"
)

type SpySleeper struct {
	Calls int
}

// We're making a mock of the real Sleep() method. Instead of it being a way to delay the program from running, it will increment the Call field of the SpySleeper struct
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// This is another mock of the Sleeper. We need a struct to hold data that we can assert in tests
type SpyCountdownOperations struct {
	Calls []string
}

// This mock of the Sleep() function that the Sleeper interface stipulates. This one will append the word "sleep" to the Call field's array
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// This mock of the Sleep() function that the Sleeper interface stipulates. This one will append the word "write" to the Call field's array
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	// since the return types are given names in the declaration using just 'return' is sufficient to implicitly return both values
	return
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
	t.Run("test output", func(t *testing.T) {
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
	})
	t.Run("test order of calls: sleep is called before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write, sleep,
			write, sleep,
			write, sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
