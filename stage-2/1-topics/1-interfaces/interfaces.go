package main

import (
	"fmt"
	"math"
)

// Personal organization structs -> interfaces -> methods > functions > main

// Starting Struct Definitions
type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
}

// Square implemented just to be a showcase of an unknown shape in the switch type assertion
type Square struct {
	width  float64
	length float64
}

// Starting Interface Definitions
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Starting Method Definitions
func (c Circle) Area() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (c Circle) Perimeter() float64 {
	return 2 * (c.radius * math.Pi)
}

func (c Circle) Diameter() float64 {
	return 2 * c.radius
}

func (r Rectangle) Area() float64 {
	return r.width * r.length
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.length)
}

func (s Square) Area() float64 {
	panic("not implemented")
}

// uncomment to work the interface for Square
//func (s Square) Perimeter() float64 {
//	panic("not implemented")
//}

// Function Definitions
func CalculateTotalArea(s ...Shape) float64 {
	var total float64
	for _, shape := range s {
		total += shape.Area()
	}
	return total
}

func detectShapes(s Shape) {
	// comma ok idiom type assertion

	//if r, ok := s.(Rectangle); ok {
	//	fmt.Println("rectangle detected", r)
	//}

	//	switch type assertion
	
	switch s.(type) {
	case Rectangle:
		fmt.Println("rectangle detected")
	case Circle:
		fmt.Println("circle detected")
	default:
		fmt.Println("unknown shape")
	}
}

func main() {
	r := Rectangle{length: 10, width: 5}
	c := Circle{radius: 10}
	s := Square{width: 10, length: 10}
	fmt.Println("rectangle area:", r.Area())
	fmt.Println("circle area", c.Area())

	fmt.Println("Total area of shapes: ", CalculateTotalArea(r, c))

	fmt.Printf("Type of rect: %T\n", r)

	detectShapes(r)
	detectShapes(c)
	detectShapes(s) //if not all methods are implemented, the error below will be thrown
	//Cannot use s (type Square) as the type Shape
	//Type does not implement Shape as some methods are missing:
	//Perimeter() float64
}
