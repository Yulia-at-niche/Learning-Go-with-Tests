package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func Countdown(out io.Writer) {
	countdown := 3
	finalWord := "Go!"
	// while i is greater than 0, i will decrement by 1
	// i is 3. First time i is 3, then 2, then 1 at which point the loop will decrement it to 0 and it won't run anymore.
	for i := countdown; i > 0; i-- {
		// This print function automatically appends a newline
		fmt.Fprintln(out, i)
		// Put a pause between each number
		time.Sleep(1 * time.Second)
	}
	// This won't append a newline
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
