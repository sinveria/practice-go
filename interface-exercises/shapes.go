package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Square: %.2f\n", s.Area())
	fmt.Printf("Perimeter/Circumference: %.2f\n\n", s.Perimeter())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	fmt.Println("Rectangle:")
	PrintShapeInfo(rect)

	fmt.Println("Circle:")
	PrintShapeInfo(circle)

	shapes := []Shape{rect, circle}
	fmt.Println("All forms:")
	for _, shape := range shapes {
		PrintShapeInfo(shape)
	}
}
