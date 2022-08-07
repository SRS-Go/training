package main

import "fmt"

func main() {
	//switch case
	ans := 10
	switch ans {
	case 1, -1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	case 6:
		fmt.Println("Six")
	case 7:
		fmt.Println("Seven")
	case 8:
		fmt.Println("Eight")
	default:
		fmt.Println("Not found")
	}
}
