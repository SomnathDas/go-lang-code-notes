package main

import (
	"fmt"
)

func pointers() {
	fmt.Println("Phase 9, Let's Go!")

	// Literal copying
	a := 10
	b := 20
	fmt.Printf("Before swap, a is %v and b is %v\n", a, b)
	b = a
	a = 50
	fmt.Printf("After swap, a is %v and b is %v\n", a, b)

	// Above, a is changed but b is still same because swap above there has literally copied the value
	// To deviate from this behaviour, we use pointers

	// Referencing
	// * is referencing/de-referencing operator
	// & is used to point to address of the further variable
	// &username <- this is saying that "point to the memory address of username variable"
	var c int = 10
	var d *int = &c // "*" here is declaring a pointer to a data type of int
	fmt.Printf("c is %v and d is %v\n", c, d)
	// to observe the actual value located in the given memory address
	// use *variableName to de-reference it
	fmt.Printf("c is %v and d is %v\n", c, *d)

	// now chaning value of c will affect d since d is a pointer that is pointing to the memory address of the value which c has
	// Go Lang does not allow Pointer arithmatic by design. We can refer to inbuilt package called "unsafe" if deem neccessary

	type User struct {
		username string
	}

	var myUser *User
	myUser = &User{username: "Uno"}
	fmt.Println("myUser struct pointer ", myUser)

	// setting and getting properties in a struct pointer
	(*myUser).username = "Panam"
	// () is there because .(dot) takes priority than * (asterisk)
	// *myUser -> de-referencing the myUser pointer variable then selecting username prop and assigning the value
	fmt.Println("get myUser.username: ", (*myUser).username)
}
