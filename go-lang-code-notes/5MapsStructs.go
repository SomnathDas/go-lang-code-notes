package main

import (
	"fmt"
)

func mapsStructs() {
	fmt.Println("Phase 5, Let's Go!")

	// Maps
	// Literal Syntax of creating a map
	// map[type of key]type of value
	repairCosts := map[string]int32{
		"Gear":   200,
		"Chasis": 450,
		"Tyre":   100,
	}

	fmt.Println(repairCosts)

	// Using make function
	corpoLevel := make(map[string]int32)
	corpoLevel = map[string]int32{
		"Arasaka":  1245,
		"Militech": 1246,
		"Kang Tao": 1230,
	}

	fmt.Println(corpoLevel)

	customerIDs := map[int32]string{
		12452:  "Jessica Jones",
		825662: "Matt Murdock",
		99022:  "Peter Parker",
	}

	fmt.Println(customerIDs)

	// Maps are reference type meaning manipulating one variable pointing to a map is going to affect all variables pointing to the same map
	myCustomerIDs := customerIDs // myCustomerIDs points to the memory location of customerIDs, no literal copying occurs
	myCustomerIDs[12452] = "Daniel Rand"
	fmt.Println(myCustomerIDs)
	fmt.Println(customerIDs)

	// In Maps, ORDER OF KEY,VALUE PAIRS (ELEMENTS) ARE NOT FIXED UNLIKE IN ARRAY OR SLICE
	myShoppingList := map[string]int32{
		"Cookies":     10,
		"Bananas":     7,
		"Cold Drinks": 5,
		"Melons":      2,
	}

	fmt.Printf("myShopping List (Before addition) : %v\n", myShoppingList)
	// Adding another entry
	myShoppingList["Bread"] = 2
	fmt.Printf("myShopping List (After addition) : %v\n", myShoppingList)

	// entry can be deleted like following
	delete(myShoppingList, "Melons")
	fmt.Println(myShoppingList)

	// Recommend way to use MAPS to make sure we don't get the result (i.e to say 0) for a key that doesn't exist
	quantity, ok := myShoppingList["Tomato"]
	fmt.Println(quantity, ok) // here quantity = 0 but ok = false which means key is not found in our map

	// len() function works for MAPS too

	// structs
	type User struct {
		username string
		userId   int32
		cart     []int32
	}

	myUser := User{
		username: "rapidfire",
		userId:   7644,
		cart: []int32{
			14555, 98744, 5664, 2344,
		},
	}

	fmt.Println(myUser.cart)

	// structs behave same as arrays when it comes to copying, it literally copies instead of referencing and hence to avoid this behaviour use pointers
	newUser := User{
		username: "silvericon",
		userId:   3411,
		cart: []int32{
			1233, 5152, 12990, 44902,
		},
	}

	newUser1 := newUser
	newUser1.username = "goldenicon"

	fmt.Println("newUser ", newUser)
	fmt.Println("newUser1 ", newUser1)

	newUser2 := &newUser              // referencing newUser instead of copying
	newUser2.username = "diamondicon" // changes username property in both newUser and newUser2

	fmt.Println("newUser ", newUser)
	fmt.Println("newUser2 ", newUser2)

	// there is no inherentence while using structs but we can use composition
	// we can use composition by doing what's called EMBEDDING

	type animal struct {
		name   string
		origin string
	}

	type bird struct {
		animal // Embedding Animal struct within Bird Struct
		canFly bool
	}

	// Struct Embedding
	emu := bird{
		animal: animal{name: "Emu",
			origin: "Australia"},
		canFly: false,
	}

	fmt.Println(emu)

}
