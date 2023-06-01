package main

import (
	"fmt"
)

func dataTypes() {
	fmt.Println("Phase 2. Go!")

	var isOpen bool = false // boolean in Go lang is its own type, it is not an alias or int i.e to say 0 != false and 1 != true
	var isSafe bool
	isSafe = true
	isWorking := true

	fmt.Printf("%v of the type: %T \n", isWorking, isWorking)
	fmt.Printf("%v of the type: %T \n", isOpen, isOpen)
	fmt.Printf("%v of the type: %T \n", isSafe, isSafe)

	var signedInt int32 = -124 // works as you would expect, stores both -ve and +ve value
	var unsignedInt uint32 = 0 // stores 0 and above positive values

	fmt.Printf("%v of the type: %T \n", signedInt, signedInt)
	fmt.Printf("%v of the type: %T \n", unsignedInt, unsignedInt)

	n := 16             // 2^4
	fmt.Println(n << 3) // 2^4 * 2^3 == 2^7 == 128
	fmt.Println(n >> 4) // 2^4 / 2^4 == 2^0 == 1

	var complexNum complex64 = 2 + 4i // Complex num is treated as a First class citizen in Go lang
	var imaginaryNum complex64 = 6i
	var realNum complex64 = 1
	fmt.Printf("%v of the type: %T \n", complexNum, complexNum)
	fmt.Printf("%v of the type: %T \n", imaginaryNum, imaginaryNum)
	fmt.Printf("%v of the type: %T \n", realNum, realNum)

	fmt.Printf("%v is the real part and %v is the imaginary part \n", real(complexNum), imag(complexNum))

	fmt.Printf("real part of type: %T and imaginary part of type: %T \n", real(complexNum), imag(complexNum))

	var altComplex complex128 = complex(24, 6)

	fmt.Printf("%v is the real part and %v is the imaginary part \n", real(altComplex), imag(altComplex))

}
