// Package raindrops converts an integer number into string phrases
package raindrops

import (
	"strconv"
	"strings"
)

// Convert function takes in a number of type Uint
// checks for the factors of the number and returns
// "Pling" if it s 3, "Plang" if it is 5, and "Plong" if it is 7
func Convert(number int) string {

	var results []string

	// Checking for factors in the given number
	if number%3 == 0 {
		results = append(results, "Pling")
	}
	if number%5 == 0 {
		results = append(results, "Plang")
	}
	if number%7 == 0 {
		results = append(results, "Plong")
	}
	if len(results) == 0 {
		// Converts the number into a strng
		numberAsString := strconv.Itoa(number)
		return numberAsString
	}
	// aggregates all contents of the slice into one word eliminating
	// the spaces between words
	rainDropSounds := strings.Join(results, "")
	return rainDropSounds
}
