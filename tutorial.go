package main

import "fmt"

/*Tutorial for data types, fmt module */
func main() {
	var number uint16 = 260 /*define a variable*/
	number = number + 5
	var number2 = 123 /*compliler determines the type*/
	number3 := 300    /*walarus operator, tell compiler to determine the type*/
	var bl bool       /*defaults for variables*/
	fmt.Println("Hello world!", number, number2, number3, bl)

	fmt.Printf("%T %T %T %T", number, number2, number3, bl)

	fmt.Printf("\nNumber is %b\t%o\t%d\t%x", 1234, 1234, 1234, 1234) /*binary, octal, decimal and hexadecimal*/
	/*output - Number is 10011010010   2322    1234    4d2*/
	var x = 3.4678901928
	fmt.Printf("\nFloating point representation %f, %F, %e, %g", x, x, x, x) /*Floating point representations*/
	/*output - Floating point representation 3.467890, 3.467890, 3.467890e+00, 3.4678901928*/

	var text = "Sample"
	fmt.Printf("\nstring is %s, double quoted version %q", text, text)
	/*string is Sample, double quoted version "Sample"*/

	fmt.Printf("\nThis is how we print %f%%", x)
	/*This is how we print 3.467890%*/

	/*Width, Precision, Padding*/
	fmt.Printf("\nText with default width and default precision %s", text)                /*Text with default width and default precision Sample*/
	fmt.Printf("\nText with width 9 and default precision is %9s", text)                  /*Text with width 9 and default precision is    Sample*/
	fmt.Printf("\nText with width -9 and default precision is %-9s negative width", text) /*Text with width -9 and default precision is Sample    negative width*/
	fmt.Printf("\nNumber with precision 2 is %.2f", x)                                    /*Number with precision 2 is 3.47*/
	fmt.Printf("\nNumber padding example: %07d", 44)                                      /*Number padding example: 0000044*/

	var out string = fmt.Sprintf("Number: %d, padded number: %07d", 44, 44) /*Save the format to a variable*/
	fmt.Print(out)                                                          /*Number padding example: 0000044Number: 44, padded number: 0000044*/
}
