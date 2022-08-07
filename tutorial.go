package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 5, 6, 7, 4, 11}
	fmt.Println(a)
	for i, element := range a {
		fmt.Println(i, element)
	}

	// if not using index
	for _, element := range a {
		fmt.Println(element)
	}

	// print only unique element
	for i, element := range a {
		for j := i + 1; j < len(a); j++ {
			if a[j] == element {
				fmt.Println("duplicates:", element)
			}
		}
	}
	// print 2d array
	var arr2d [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, element := range arr2d {
		for _, element2 := range element {
			fmt.Println(element2)
		}
	}
}
