package main

//this is how you import multiple packages
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*Tutorial for input data, change type */
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the year you were born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("you will be %d years old at the end of 2022.", 2022-input)
}
