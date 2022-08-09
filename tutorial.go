package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rect struct {
	width  float64
	length float64
}

func (r *Rect) area() float64 {
	return r.width * r.length
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c1 := Circle{3.4}
	r1 := Rect{4.1, 2}
	shapes := []Shape{&c1, &r1}

	for _, shape := range shapes {
		fmt.Println("using a wrapper function", getArea(shape))
		fmt.Println("Access through interface", shape.area())
	}
}
