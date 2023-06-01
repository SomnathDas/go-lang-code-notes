package main

import (
	"fmt"
)

func arraySlices() {
	fmt.Println("Phase 4, Let's Go!")

	grades := [4]int{25, 25, 24, 23} // Our Array
	fmt.Printf("%v of the type: %T\n", grades, grades)

	var students [3]string
	students[0] = "Arasaka"
	students[2] = "Kang Tao"
	students[1] = "Militech"

	fmt.Printf("%v of the type: %T\n", students, students)

	identityMatrix := [3][3]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}} // Array of array
	fmt.Println(identityMatrix)

	// Copying arrays in Go lang
	a := [3]int32{1, 2, 3}
	b := a //  In go lang, "b" does not point to the same location in memory as "a", go lang creates a literal copy of a
	// above can be proven by the following:
	b[1] = 5 // Changing b should not mutate a
	// we can observe that by
	fmt.Printf("Array a : %v\n", a)
	fmt.Printf("Array b : %v\n", b)

	// if we do not want this behaviour then we can use "Pointers"
	c := &a // here a is an array in itself stored somewhere in the memory and c is pointing to a's location in the memory
	// so if we manipulate c, it should alter a since c is pointing to a
	c[2] = 40
	fmt.Printf("Array a : %v\n", a)
	fmt.Printf("Array c : %v\n", c)

	// Slices; Slices are reference type by foundation
	var mySlice = []int32{1, 2, 3} // Our slice
	fmt.Printf("Slice %v of the type: %T\n", mySlice, mySlice)

	// as mentioned earlier Slices are reference type
	// we can observe this behaviour with the following illustration
	myOtherSlice := mySlice
	myOtherSlice[2] = 69

	fmt.Printf("mySlice: %v\n", mySlice)
	fmt.Printf("myOtherSlice: %v\n", myOtherSlice)
	// Whereas in Array, Go lang was making a literal copy of the array
	// So this is one major difference b/w Arrays and Slices

	fmt.Println("Length of the students Array : ", len(students))
	fmt.Println("Length of the mySlice Slice : ", len(mySlice))

	fmt.Println("Capacity of the mySlice Slice : ", cap(mySlice))

	// Lets slice up the slices
	g := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	h := g[:]   // Slice all elements of g
	i := g[2:]  // Slice elements from 3rd position in the array (which starts from index 2) till the end
	j := g[2:4] // Slice elements from 3rd position (here which is "3") in the array up till 4th position (here which is "4")
	k := g[:6]  // Slice elements from beginning up till 7th position (6th index)

	// In summary, starting index is inclusive and ending index is exclusive while slicing

	fmt.Println("Array h: ", h)
	fmt.Println("Array i: ", i)
	fmt.Println("Array j: ", j)
	fmt.Println("Array k: ", k)

	// Slicing also works with arrays
	l := [...]int32{1, 2, 3, 4} // The compiler can identify the length of an array, based on the elements specified in the array declaration.
	m := l[1:2]
	fmt.Println("Array m: ", m)

	// Another way to create a slice is by using make function which also can be used for other things
	altSlice := make([]int32, 4, 100) // slice of type int32 whose length is 4 and capacity is 100
	fmt.Println("altSlice Slice: ", altSlice)
	fmt.Println("Length of altSlice Slice: ", len(altSlice))
	fmt.Println("Capacity of altSlice Slice: ", cap(altSlice))

	// Concatenating two slices
	alpha := []int32{1, 2}
	beta := []int32{99, 98}
	gamma := append(alpha, beta...) // Kind of like javascript spread operator
	fmt.Println(gamma)              // 1,2,99,98
	// append function may cause expensive copying operation if the underlying array is too large

}
