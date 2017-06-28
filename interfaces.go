// Interfaces.go
// Interfaces are named collections of method signatures.
// From https://play.golang.org/p/4kN4P9C2AW
// Explained in https://gobyexample.com/interfaces
// and http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import (
	"fmt"
	"math"
)

// An interface is two things: it is a set of methods, but it is also a type.
// CONVENTION: No blank line between a type and its associated func.

type square struct {
	side float64
}
func (s square) area() float64 {
	return s.side * s.side
}

// In this example, area() appears in several funcs.
// area() is a method associated with different types.
type circle struct {
	radius float64
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}
func printArea(x shape) {
	fmt.Println("My area is:", x.area())
}


// Other code examples: https://play.golang.org/p/yGTd4MtgD5
type Animal interface {
	Speak() string
}

type Dog struct {
}
func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}
func (c Cat) Speak() string {
	return "Meow!"
}

type Llama struct {
}
func (l Llama) Speak() string {
	return "?????"
}

type JavaProgrammer struct {
}
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {  // start figuring out a program from the main

	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	  /* Output:
		//Woof!
    //Meow!
    //?????
    // Design patterns!
    */
	}

	s1 := square{10}
	c1 := circle{4}   // shape c1
	printArea(s1)
	printArea(c1)  // given a shape
  // Output: My area is: 100
  // My area is: 50.26548245743669
}
