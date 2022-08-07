package main

//this is how you import multiple packages
import (
	"fmt"
)

func main() {
	//condtional and boolean expressions
	// operators < > <= >= == !=
	x := 5
	y := 6
	val := x < y
	fmt.Printf("%t", val)
	val = x == y
	fmt.Printf("\n%t", val)
}
