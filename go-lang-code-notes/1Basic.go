package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func basic() {
	var number float32 = 23.2 // Common way of declaring variables and initialising variables
	fmt.Printf("Type of number is: %T \n", number)

	number2 := 32 // Standard way of declaring and initialising variables
	fmt.Printf("Type of number2 is: %T \n", number2)

	fmt.Println("Enter your name :> ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	fmt.Printf("Type of reader is: %T \n", reader)
	fmt.Printf("Type of input is: %T \n", input)
	fmt.Printf("Type of err is: %T \n", err)

	var txt = "Hello World"
	fmt.Printf("%v of type: %T", txt, txt) // All Fine

	text := "Hello World"
	fmt.Printf("%T \n", text) // All Fine

	// Conversions

	num1 := 42
	fmt.Printf("%v of type: %T \n", num1, num1)
	fmt.Printf("%v of type : %T \n", string(num1), string(num1)) // Expected "42" ? Nope, this type of conversion takes it to be ASCII Code "*"
	// use strconv package for num to string conversion
	fmt.Printf("%v of type : %T \n", strconv.Itoa(num1), strconv.Itoa(num1))

	// Go lang does not recommend doing implicit conversion such as that of between float32 or float64 to integer
	// it can lead to unaware loss of information

}

func alt() {
	// Figured out the type above
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Println(input)
	if err != nil {
		return
	}
}

/*

First letter capital in the package scope is used to declare variables which EXPORTS outside the package
First letter small is used to declare variables in both package and function scope to declare those which are going
to be used inside the package
There is no such thing as private scope in GO lang

*/
