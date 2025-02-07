// Are interfaces potentially a dictionary of functions?
// Maybe interfaces are like methods for objects?
// Almost looks like inheritance

package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	area() float64
}

// Define a circle
type Circle struct {
	x, y, radius float64
}

// Define a rectangle
type Rectangle struct {
	width, height float64
}

// Define a method for circle (implementation of Shape.area())
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// Define a method for rectangle (implementation of Shape.area())
func (rect Rectangle) area() float64 {
	return rect.width * rect.height
}

// Define a method for shape
func get_area(shape Shape) float64 {
	return shape.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle area: %f\n", get_area(circle))
	fmt.Printf("Rectangle area: %f\n", get_area(rectangle))
}
