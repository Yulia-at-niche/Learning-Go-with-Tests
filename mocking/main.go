package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(out io.Writer) {
	countdown := 3
	// while i is greater than 0, i will decrement by 1
	// i is 3. First time i is 3, then 2, then 1 at which point the loop will decrement it to 0 and it won't run anymore.
	for i := countdown; i > 0; i-- {
		// This print function automatically appends a newline
		fmt.Fprintln(out, i)
	}
	// This won't append a newline
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
