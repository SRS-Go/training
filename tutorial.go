package main

import "fmt"

func main() {
	//for loops and while loops
	// there is no while loop in Go, for can work like while loop
	x := 0
	for x <= 5 {
		fmt.Println("value of x=", x)
		x++
	}

	for x := 0; x <= 100; x++ {

		if x == 15 {
			fmt.Printf("Skipped %d", x)
			continue
		}
		if x == 20 {
			fmt.Printf("Breaking loop at %d", x)
			break
		}
		fmt.Println("x=", x)
	}
}
