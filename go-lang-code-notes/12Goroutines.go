package main

import (
	"fmt"
	"time"
)

func goRoutines() {
	fmt.Println("Phase 12, Let's Go!")
	anonVar := "Anon Hello"
	anonVarBetter := "Anon Better Hello"

	// creating our very first goRoutine
	go sayHello()

	go func() {
		fmt.Println(anonVar)
	}()

	/*
		above, what is happening is called "Closure" in go-lang
		anonymous function up there has access to its outer scope and therefore can take advantage of anonVar variable declared above there
		and use it inside the go routine even though the go-routine is running with a completely different execution stack.
		Go runtime understands where to get anonVar variable from

		Problem arises that is we have created a dependancy between the variable in main function and variable in the go routine
	*/

	// illustration of the above said problem
	simpleMsg := "Illustration msg"
	go func() {
		fmt.Println(simpleMsg)
	}()
	simpleMsg = "Illustration magzine"

	/*
		what is happening in the above said illustration is that: we've expected "Illustration msg" to be printed out not "Illustration magzine"
		because most of the time the GO scheduler is not going to interrupt the main thread until it hits the sleep call below which means
		even though it launches the another go routine there, it still is exeucting the main function. So it actually gets from line 30 to line34
		and re-assign simpleMsg value before the go Routine has the chance to print it out
	*/

	// best practice while using anon functions
	go func(s string) { // passing by value (copying) and hence it decouples the dependancy
		fmt.Println(s)
	}(anonVarBetter)

	time.Sleep(100 * time.Millisecond) // sleep calls not recommended instead use waitGroup

	/*
		adding "go" keyword infront of a function invocation is going to tell go-lang to spin-off a "green thread"
		and run that function in that green thread
	*/

	/*
		Most programming languages use OS threads i.e they got individual function call stack dedicated to the execution
		of whatever code is handed to that thread. These tend to be very very large. Here comes Green Threads
	*/

	/*
		Green threads emulate multithreaded environments without relying on any native OS abilities,
		and they are managed in user space instead of kernel space,
		enabling them to work in environments that do not have native thread support.
	*/

}

func sayHello() {
	fmt.Println("Hello!")
}
