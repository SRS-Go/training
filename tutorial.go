package main

//this is how you import multiple packages
import "fmt"

func main() {
	//condtional statements - if else elif
	condtion := bool(true) == true
	fmt.Println("This will always print")
	if !condtion {
		fmt.Println("This is inside if")
	} else {
		fmt.Println("This is inside Else")
	}

	age := 17
	if age >= 18 {
		fmt.Println("You can ride alone!")
	} else if age >= 14 {
		fmt.Println("You can ride with your parent!")
	} else {
		fmt.Printf("wait for %d years before you ride", 18-age)
	}
}
