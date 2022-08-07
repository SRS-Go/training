package main

//this is how you import multiple packages
import "fmt"

func main() {
	//chained conditions:
	// and - &&
	// or - ||
	// not !
	x := 5
	y := 6
	z := 7
	a := x == y
	val := x < y && x < z
	fmt.Printf("%t", val)
	val = x < y || x < z
	fmt.Printf("\n%t", val)
	val = x < y && x > z
	fmt.Printf("\n%t", val)
	val = x > y || x > z
	fmt.Printf("\n%t", val)
	val = x < y || x > y && x > z //not a good way of writing code
	fmt.Printf("\n%t", val)
	val = x < y || (x > y && x > z)
	fmt.Printf("\n%t", val)
	val = x < y && !a
	fmt.Printf("\n%t", val)
}
