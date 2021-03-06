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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle : Radius = %v", c.Radius)
}

type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2*r.Height + 2*r.Width
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle : Height = %v, Width = %v", r.Height, r.Width)
}

func printArea(shape ShapeWithArea) {
	fmt.Println(shape.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func printPerimeter(shape ShapeWithPerimeter) {
	fmt.Println(shape.Perimeter())
}

type Dimension interface {
	Area() float32
	Perimeter() float32
}

/* type Dimension interface {
	ShapeWithArea
	ShapeWithPerimeter
} */

func printDimension(shape Dimension) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimeter())
}

func main() {
	circle := &Circle{Radius: 9}
	//printArea(circle)
	fmt.Println(circle)
	rectangle := &Rectangle{Height: 6, Width: 5}
	//printArea(rectangle)
	fmt.Println(rectangle)
	//printPerimeter(circle)
	//printPerimeter(rectangle)
	printDimension(circle)
	printDimension(rectangle)

}
