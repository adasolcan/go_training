// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

import "fmt"

// Declare a type named user.
type user struct {
	id   int
	name string
}

// Create a function that changes the value of one of the user fields.
func changeName(u *user, name string) {
	// Use the pointer to change the value that the
	// pointer points to.
	u.name = name
}

func main() {

	// Create a variable of type user and initialize each field.
	my_user := user{
		id:   2,
		name: "John",
	}

	// Display the value of the variable.
	fmt.Println("Value:", my_user)

	// Share the variable with the function you declared above.
	changeName(&my_user, "Jane")

	// Display the value of the variable.
	fmt.Println("Value:", my_user)
}
