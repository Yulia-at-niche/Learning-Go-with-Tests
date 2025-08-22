package structs

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

/*
A method is a function with a receiver. A method declaration binds an identifier, the method name, to a method, and associates the method with the receiver's base type.

Declaration:
func (r Rectangle) Area() float64

Identifier i.e method's name:
Area is the identifier aka the method's name.
The method's arguments are anything in the '()' of Area(). In this case there are none.

The method that the identifier is binding:
Whatever is in the '{}' i.e the body. The actual logic this method will do

Return type:
the 'float64' of the declaration

The receiver:
The (r Rectangle) in the declaration.
The 'r' is the receiver variable. We'll be able to use this in our method.
It is a convention in Go to have the receiver variable be the first letter of the type.

Receiver's base type:
Rectangle in the declaration

To put it all together:
func (receiverVar ReceiverType) identifier(argsIfAny) returnType { // End of the declaration
	// ...Begin the method logic that binds to the identifier
}
*/

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
