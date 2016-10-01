// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {

	// Declare an array of 5 strings set to its zero value.
	var names1 [5]string

	// Declare an array of 5 strings and pre-populate it with names.
	names2 := [5]string{"John1", "John2", "Joh3", "John4", "John5"}

	// Assign the populated array to the array of zero values.
	names1 = names2

	// Iterate over the first array declared.
	for i, name := range names1 {
		// Display the string value and address of each element.
		fmt.Println("Value ", name, " Address ", &names1[i])
	}
}
