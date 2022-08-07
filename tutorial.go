package main

import "fmt"

func main() {
	var x [6]int = [6]int{1, 2, 3, 4, 5, 6}
	var s []int = x[2:5] //where to start slice and where to end (excluding the end)
	fmt.Println(x)
	fmt.Println(s, len(s), cap(s)) //length gives the slice lenght, while cap gives the initial array capacity
	fmt.Println(s[:cap(s)])

	var a []int = []int{11, 22, 33, 44, 55} //Cheat on Go to make array dynamic, initialize with no length
	fmt.Println(cap(a[:3]))

	a = append(a, 66) //use append method to append element and store it as initial array name
	fmt.Println(a)

	b := make([]int, 5) //create empty slice
	fmt.Println(b)
}
