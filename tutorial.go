package main

import "fmt"

func changeValue(str *string) {
	*str = "changed!"
}

func changeValue2(str string) {
	str = "changed again!"
}

func main() {
	x := 7
	y := &x
	fmt.Println(x, y)
	*y = 8
	fmt.Println(x, y)

	toChange := "hello"
	changeValue(&toChange) // pass pointer if you want to change value
	fmt.Println(toChange)
	changeValue2(toChange)
	fmt.Println(toChange)
}
