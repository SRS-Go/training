package main

import "fmt"

func main() {
	//array in Go
	var arr [5]int
	fmt.Println(arr) //[0 0 0 0 0]
	arr[3] = 6
	fmt.Println(arr)
	//arr[6] = 6 // index out of range error
	//fmt.Println(arr)

	arr1 := [3]int{4, 5, 6}
	sum := 0
	fmt.Println(arr1)
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i]
	}
	fmt.Println(sum)

	//multi dimetional array
	arr2d := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr2d)
	fmt.Println(arr2d[0][2]) // access elements
}
