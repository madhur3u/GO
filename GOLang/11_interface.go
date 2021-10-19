package main

import (
	"fmt"
	"math"
)

// we created an INTERFACE shape
// this has 2 methods
// we can call these two methods with this interface
// without using the types
type Shape interface {
	Area() float64
	Perimeter() float64
}

// in this we created two types of circle and rectangle
// and both have 2 methoda area and perimeter
// here we will create an INTERFACE which will provide us
// area and perimeter and we will not diectly deal with these types
// so interface will be used to use these types
type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

func (c Circle) Area() float64 {
	return (math.Pi * c.radius * c.radius)
}

func (c Circle) Perimeter() float64 {
	return (math.Pi * 2 * c.radius)
}

func (r Rectangle) Area() float64 {
	return (r.width * r.length)
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {

	fmt.Println("Hello World")

	// we made a object shapes of our interface shape
	// it is a slice and has content as our types we made
	// we can access these just using shapes[i]
	shapes := []Shape{Circle{5.2}, Rectangle{2, 8}}

	// now using this interface, we can call both methods
	// but using interface we cannot use shape.radius or shape.length
	// as shape only has 2 things area and perimeter in it
	fmt.Println(shapes[0].Area())
	fmt.Println(shapes[0].Perimeter())
	fmt.Println(shapes[1].Area())
	fmt.Println(shapes[1].Perimeter())

}
