package main

import "fmt"

//this is how you define struct
type Point struct {
	x int32
	y int32
	z bool
}

func changeX(pt *Point) {
	pt.x = 100
}

type Circle struct {
	radius float64
	center *Point
}

type AnotherCircle struct {
	radius float64
	*Point // no need to give a name.. as long as there are no common attributes on both structs it will work
}

func main() {
	var p1 Point = Point{1, 2, false}
	fmt.Println(p1.x, p1.y)
	var p2 Point = Point{x: 2} // passing one value, defaults other
	fmt.Println(p2.x, p2.y)
	p3 := &Point{y: 3} // assign as pointer instead of object, benefits if you want to change values in other functions
	changeX(p3)
	fmt.Println(p3.x)
	// struct inside a struct
	p4 := &Point{x: 7}
	c1 := &Circle{3.75, p4}
	fmt.Println(c1)
	c2 := &Circle{3.75, &Point{4, 5, true}}
	fmt.Println(c2, c2.center, c2.center.y)
	c3 := &AnotherCircle{3.4, &Point{x: 4}}
	fmt.Println(c3.radius, c3.x, c3.y, c3.z) // access the inner struct attribute directly
}
