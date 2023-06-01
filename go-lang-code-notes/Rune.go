package main

import "fmt"

func runeFunc() {
	// rune -> is an alias for the type of int32 in go lang
	// rune -> is used to store UNICODE CODE-POINTS ex. U+03C0
	// since int32 -> 32bits -> 4 bytes -> meaning int32 is capable of storing up to 4 bytes
	// unicode code-points at max requires 4 bytes to be stored
	// hence Rune is used to store unicode code points

	myPieRune := 'π'
	fmt.Printf("%U\n", myPieRune) // 'U' format flag is used to show unicode code-point
	fmt.Printf("%c\n", myPieRune) // 'c' format flag is used to show unicode character

	mySaluteRune := '🫡'
	fmt.Printf("%U\n", mySaluteRune) // 'U' format flag is used to show unicode code-point
	fmt.Printf("%c\n", mySaluteRune) // 'c' format flag is used to show unicode character

	/*
		Since the rune type represents a Unicode character, a string in Go is often referred to as a sequence of runes.
		However, runes are stored as 1, 2, 3, or 4 bytes depending on the character.
		Due to this, strings are really just a sequence of bytes.
		In Go, slices are used to represent sequences and these slices can be iterated over using range.
	*/

}
