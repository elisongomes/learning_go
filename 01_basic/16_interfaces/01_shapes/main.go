package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

func typeArea(s Shape) {
	fmt.Printf("The area of shape is %0.2f\n", s.area())
}

type Retangle struct {
	height float64
	width float64
}

func (r Retangle) area() float64  {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64  {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	fmt.Println("Interfaces")

	retangle1 := Retangle{10, 40}
	typeArea(retangle1)

	circle1 := Circle{25}
	typeArea(circle1)
}
