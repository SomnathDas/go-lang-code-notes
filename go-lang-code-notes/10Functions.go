package main

import (
	"fmt"
)

func functions() {
	fmt.Println("Phase 10, Let's Go!")

	multiArg("sd", "un@usa.com")

	currentCodeToPass := 10 // parent int variable set to 10
	println("currentCodeToPass ", currentCodeToPass)
	passByRef(&currentCodeToPass)                    // passing the argument by reference instead of literal value
	println("currentCodeToPass ", currentCodeToPass) // changed parent variable to 20

	// func with variadic params call
	sumOfNums("Total sum is:", 1, 2, 3, 4, 5)

	// using divide func with error handling and multiple return values
	/*quotient, err := divide(12.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quotient)
	*/

	//counter, structs and methods
	myUser := User{username: "samurai"}
	fmt.Println("username of myUser:", myUser.getUsername())

	// Structs and Methods with Pointers to manipulate underlying data
	myCounter := Counter{countValue: 0}
	myCounter.count()
	myCounter.count()
	fmt.Println("myCounter countValue is:", myCounter.countValue)

}

// passing multiple args of same data type
func multiArg(username, email string) bool {
	return false
}

// default behaviour of args of a function is that they are passed by value
// to pass by reference
func passByRef(currentCode *int) {
	// de-referencing and re-assigning new value that will not only change the value of currentCode in local context but also in where it is called
	(*currentCode) = 20
	fmt.Println("passByRef done")
}

// variadic parameters
// variadic params must always be at last of arguments
func sumOfNums(msg string, num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	fmt.Println(msg, sum)
	return sum
}

// Multiple return values and Error handling Sample
func divide(num1, num2 float64) (float64, error) {
	if num2 == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return num1 / num2, nil
}

// Structs and Methods
type User struct {
	username string
}

func (userObj User) getUsername() string {
	return userObj.username
}

// Structs and Methods with Pointers to manipulate underlying data
type Counter struct {
	countValue int
}

func (counterObj *Counter) count() {
	(*counterObj).countValue += 1
}
