package main

import (
	"fmt"
)

func interfaces() {
	fmt.Println("Phase 11, Let's Go!")

	var mySquare Shape = Square{
		width:  10,
		length: 20,
	}
	fmt.Println("Area of mySquare is:", mySquare.Area())

	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello, World!"))

	counter := IntCounter(0)
	counter.Increment()

}

// an interface defines method signatures which are then implmeneted by the type (structs)
// method signatures are methodName, methodArguments, methodReturnType
type Shape interface {
	Area() float32
}

type Square struct {
	length float32
	width  float32
}

// implmeneted by the type (structs)
func (sObj Square) Area() float32 {
	return sObj.length * sObj.width
}

// Interfaces do not describe data like struct, they describe behaviour
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

// we do not have to use structs to implement interfaces
type Incrementer interface {
	Increment() int
}

type IntCounter int

func (ic *IntCounter) Increment() int {
	*ic = *ic + 1
	return int(*ic)
}
