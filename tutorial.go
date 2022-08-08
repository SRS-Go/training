package main

import "fmt"

func test(x int) string {
	return fmt.Sprintf("hello! %d", x)
}

func test4(myFunc func(int) string) {
	fmt.Println("Inside test4, result=", myFunc(7))
}

func returnFunc(x string) func() {
	return func() { fmt.Println(x) }
}

func main() {
	x := test // assign a value to variable and call the variable as function
	fmt.Println("calling x as test, result=", x(5))

	// this is how to create a function as a variable and call it
	test2 := func(x int) int {
		return x * -1
	}(8)
	fmt.Println("test2 called, result=", test2)

	test4(x) //here in test4 we are passing x which is a variable for func object test, test will return a string which will be printed in test4

	returnFunc("hello")()
	y := returnFunc("Hi")
	y()
}
