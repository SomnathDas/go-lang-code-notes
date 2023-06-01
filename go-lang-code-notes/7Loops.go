package main

import (
	"fmt"
)

func loops() {
	fmt.Println("Phase 7, Let's Go!")

	// for loop standard
	for i := 0; i < 11; i++ {
		fmt.Println(i)
	}

	// Multi in a single statement
	// increment that is i++ in Go lang is not an expression but a statement itself
	for i, j := 0, 0; i < 11; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// another way of working with for loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	// Or
	j := 0
	for ; j < 5; j++ {
		fmt.Println(j)
	}

	// we have continue keyword which can sometimes be very useful in a loop
	for i := 0; i < 11; i++ {
		if i%2 == 0 {
			continue // doesn't executes further block
		}
		fmt.Printf("Loop with continue keyword: %v\n", i)
	}

	// we also have break and break [enter label for the loop here]
	// break keyword will stop the nearest loop

	for i := 0; i < 16; i++ {
		if i > 6 {
			break // will simply break this loop
		}
	}

Loop:
	for i := 0; i < 16; i++ {
		for j := 0; j < 4; j++ {
			break Loop // will break the loop from the point which is mentioned using Loop label on line 50
		}
	}

	slicePets := []string{"Dog", "Cat", "Wolf", "Shark"}
	// Standard way to operate loop over arrays and slices
	for key, value := range slicePets {
		fmt.Println(key, value)
	}

	for _, value := range slicePets {
		fmt.Println(value)
	}

}
