package main

import "fmt"

// Two different interfaces
type Speaker interface {
	Speak()
}

type Walker interface {
	Walk()
}

// A struct that implements both interfaces
type Person struct {
	name string
}

// This WILL satisfy the Speaker interface because the name matches exactly
func (p *Person) Speak() {
	fmt.Printf("%s says hello!\n", p.name)
}

// You can have additional methods with different names
func (p *Person) Talk() {
	fmt.Printf("%s is talking!\n", p.name)
}

func (p *Person) Walk() {
	fmt.Printf("%s is walking\n", p.name)
}

// Function that takes both interfaces
func DoActions(speaker Speaker, walker Walker) {
	speaker.Speak() // Calls Speak() method
	walker.Walk()   // Calls Walk() method
}

// func main() {
// 	person := &Person{name: "Alice"}

// 	// Same object, but Go calls different methods based on interface type
// 	DoActions(person, person)

// 	// You can also call the concrete method directly
// 	person.Talk()
// }
