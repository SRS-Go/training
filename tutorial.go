package main

import "fmt"

//retunvalues with variable names created
func add(x int, y int) (z1 float64, z2 float64) {
	defer fmt.Println("This should happen at the end")
	z1 = float64(x) / float64(y)
	z2 = float64(y) / float64(x)
	fmt.Printf("printing values of z1=%f z2=%f\n", z1, z2)
	return z1, z2
}

//following is valid, define data type only once if same data type
func add2(x, y int) {
	fmt.Println(x + y)
}

func main() {
	x, y := add(15, 7)
	fmt.Println(x, y)
	add2(5, 6)
}
