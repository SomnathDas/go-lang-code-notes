package main

import (
	"fmt"
)

// Enumerated constants
const ( // iota starts at 0 in each block and then increments by 1
	a = iota // 0
	b = iota // 1
	c = iota // 2
	d = iota // 3
)

const (
	e = iota // 0
	f = iota // 1
	g = iota // 2
	h = iota // 3
)

func constants() {
	fmt.Println("Phase 3, Let's Go!")

	const myConstant int32 = 4209 // Constants need to be calculable at compile time
	/* For ex:  const pieByTwo = math.Asin(1); <-- this is not a constant since it is calculated at runtime*/

	/* PascalCase for exported constants and camelCase for internal constants */
	const bs = 43.55
	fmt.Printf("%v of the type: %T \n", myConstant, myConstant)
	fmt.Printf("%v of the type: %T \n", bs, bs)

	fmt.Printf("%v of the type: %T \n", a, a)
	fmt.Printf("%v of the type: %T \n", b, b)
	fmt.Printf("%v of the type: %T \n", c, c)
	fmt.Printf("%v of the type: %T \n", d, d)

	/* Typed constants can interoperated with same type only */
	/* While untyped constants can interoperate with similar types */

	const myTypedConstant int32 = 45
	const myUntypedConstant = 22

	var mySimilarTypeValue = 46

	fmt.Println(myTypedConstant + myConstant) // All Fine
	//fmt.Println(myTypedConstant +  mySimilarTypeValue); // Not fine, need to convert mySimilarTypeValue to int32 first

	fmt.Println(myUntypedConstant + mySimilarTypeValue) // All fine
}
