package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Define an interface for sleeper. The argument for countdown will use this and so will the tests for their mocks
type Sleeper interface {
	Sleep()
}

// This is the 'real' sleeper. It will implement the Sleeper interface
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	countdown := 3
	finalWord := "Go!"
	for i := countdown; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
