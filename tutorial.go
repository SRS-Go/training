package main

//this is how you import multiple packages
import (
	"fmt"
)

func main() {
	/*Arithmetic and match operations */
	var num1 int = 9
	var num2 int = 4
	answer := (num1 + num2) % 2
	fmt.Printf("%d", answer)
	var num3 float64 = 8
	var num4 int = 4
	answer2 := num3 / float64(num4)
	fmt.Printf("\n%f", answer2)
	answer3 := num1 % num2
	fmt.Printf("\n%d", answer3)
}
