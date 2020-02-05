package main

import (
	"fmt"
	"math"
)

// Shape - Define interface to be shared by all shape structs
type Shape interface {
	area() float64
}

// Circle - Struct
type Circle struct {
	x, y, radius float64
}

// Rectangle - Struct
type Rectangle struct {
	w, h float64
}

// Value Reciever
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Value Reciever
func (r Rectangle) area() float64 {
	return r.w * r.h
}

// getArea Getter Method for shapes
func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	// Interfaces allow creation of method signatures for various structs
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{w: 10, h: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
