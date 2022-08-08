package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func main() {

	// slices are mutables (maps as well)
	var x []int = []int{3, 4, 5}
	y := x
	y[0] = 100
	fmt.Println(x, y)

	// arrays are immutable
	var xx [2]int = [2]int{1, 2}
	yy := x
	yy[1] = 3
	fmt.Println(xx, yy)

	// change value in other function for mutable data
	var slice []int = []int{2, 5, 7}
	fmt.Println(slice)
	changeFirst(slice)
	fmt.Println(slice)
}
