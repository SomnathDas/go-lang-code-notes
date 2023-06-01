package main

import (
	"fmt"
)

func controlFlow() {
	fmt.Println("Phase 6, Let's Go!")

	if true {
		fmt.Println("Phase 6")
	}

	const answer int32 = 25
	var guess int32 = 25

	if guess < answer {
		fmt.Println("Colder")
	}
	if guess > answer {
		fmt.Println("Hotter")
	}
	if guess == answer {
		fmt.Println("Correct Guess!")
	}

	// Go lang also does shortcircuting. Let's say when it comes to evaluation of || (OR) logic gates. Given example:

	// Since this is an OR validatation, go lang knows that any one of these must be true for this block to be executed
	// since first is literally TRUE, GO lang skips the remaining checks and evalutes the block
	if true || thisFunction() || thatFunction() {
		println("Truthy Result")
	}

	// if you make first one false then you can see thisFunction gets evaluted and then the block
	if false || thisFunction() || thatFunction() {
		println("Truthy Result")
	}

	// Another thing to remember is that Go Lang will only execute the first block which satifies the given condition and will skip the remaining
	// Example
	omega := 25

	if omega > 24 {
		println("Omega is greater than 24") // this following if-else block will stop at first block since it fullfills the condition
		// Go lang will not evalute further even if other condition matches past this block
	} else if omega < 26 {
		println("Omega is less than 26")
	}

	// Switch
	// Shows similar behaviour as that of if-else
	// first block whose condition matches gets executed regardless of any match after that
	// each case is by-default break; (breaked) so no need for that keyword
	// if we want to make consecutive case execute that is to say fall we can use fallthrough keyword
	// fallthrough keyword is logic-less meaning it will let cases stated below execute regardless of whether the conditions are met
	i := 12 // change it to 69 or something to see the cringe of fallthrough keyword
	switch {
	case i > 10:
		println("i is greater than 10")
		// fallthrough // it will execute cases below if this case is true even when cases below are not true
	case i < 20:
		println("i is less than 20")
	default:
		println("i is neither greater than 10 nor less than 20")
	}

	// you can explicitly use "break" to stop furthur execution

}

func thisFunction() bool {
	fmt.Println("thisFunctionI() executed")
	return true
}

func thatFunction() bool {
	fmt.Println("thatFunction() executed")
	return false
}
