package main

import (
	"fmt"
	"math"
)

/*
	in C# & Java, class Circle implements ShapeWithArea
*/

type ShapeWithArea interface {
	Area() float32
}

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func printArea(shape ShapeWithArea) {
	fmt.Println(shape.Area())
}

func main() {
	circle := &Circle{Radius: 9}
	printArea(circle)

	rectangle := &Rectangle{Height: 6, Width: 5}
	printArea(rectangle)
}
