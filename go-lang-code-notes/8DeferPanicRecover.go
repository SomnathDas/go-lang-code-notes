package main

import (
	"fmt"
	"log"
)

func deferPanicRecover() {
	fmt.Println("Phase 8, Let's Go!")

	// defer
	// defer takes a function call not just a function. so it is our job to defer functions which is called
	// it moves the function call to the end of execution of the parent function but before it returns

	fmt.Println("Start")
	fmt.Println("Processing")
	fmt.Println("End")

	// output would be Start -> Processing -> End
	// now using defer keyword

	fmt.Println("Start")
	defer fmt.Println("Defer Processing")
	fmt.Println("End")

	// deferred function call when given an arugment, takes the initial state of that argument and does not bother if it eventually changes
	// deffered functions are called in the order of last-in > first-out

	myNumber := 12
	defer fmt.Println("myNumber ", myNumber) // prints 12 only not 69
	myNumber = 69

	// panic
	// panic is used instead of error only when the program goes into a state in which it cannot recover from
	// for ex. divide by zero runtime error
	// when is func is called : func execution -> then deferred func (if any) exeuction -> then panic (if any)
	panicking()
	fmt.Println("main() func execution ends")
}

func panicking() int32 {
	fmt.Println("Executing alright")
	fmt.Println("About to panic now")
	// recover
	// if we do not want parent function to end abruptly and handle the error ourselves
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
			/* if you can't deal with that error, just repanic again*/
			// panic(err)
		}
	}()
	panic("Something bad has happened")
}
