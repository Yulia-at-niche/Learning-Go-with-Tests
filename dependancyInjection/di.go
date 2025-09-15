package main

import (
	"fmt"
	"io"
	"os"
)

/*
Core Problem:
This is difficult to test because it prints to the terminal and the test can't capture that.

	func Greet(name string) {
	    fmt.Printf("Hello, %s", name)
	}

Key Insight:
What if we could be more general about how it's printed?

Research:
Both Printf uses Fprinf under the hood. Fprintf just needs an io.Writer. Any type in the Go standard library with a Write() method implements the io.Writer interface.

Solution:
So if we pass in an io.Writer to Greet(), it can become more flexible to how it prints. Because of that flexibility we can see what it's going to print by giving Greet() a buffer. We can test the content of the buffer to see if it's what we expect
*/
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Narla")
}
