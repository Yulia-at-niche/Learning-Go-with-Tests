package main

import "fmt"

//  It is good to separate your "domain" code from the outside world (side-effects). The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.

func Hello() string { // The 'String' here means that this returns a string
	return "Hello, world!" // Domain logic
}

func main() {
	fmt.Println(Hello()) // Side effect
}
