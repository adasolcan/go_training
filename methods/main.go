// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.
type ballPlayer struct {
	name         string
	atBats, hits int
}

// Declare a method that calculates the batting average for a player.
func (p ballPlayer) average() float64 {
	return float64(p.atBats) / float64(p.hits)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := [3]ballPlayer{
		{"john", 10, 3},
		{"jim", 25, 7},
		{"bill", 14, 5},
	}

	// Display the batting average for each player in the slice.
	for _, p := range players {
		fmt.Printf("%s %v\n", p.name, p.average())
	}
}
