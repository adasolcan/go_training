// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import "fmt"

type user struct {
	id   int
	name string
}

func main() {
	// Declare variable of type user and init using a struct literal.
	u1 := user{
		id:   12,
		name: "Jim",
	}

	// Display the field values.
	fmt.Printf("%v\n", u1)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		id   int
		name string
	}{
		id:   10,
		name: "John",
	}

	// Display the field values.
	fmt.Printf("%v\n", u2)
}
