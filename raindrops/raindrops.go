package raindrops

import (
	"strconv"
	"strings"
)

// Convert function takes in a number of type Uint
// checks for the factors of the number and returns
// "Pling" if it s 3, "Plang" if it is 5, and "Plong" if it is 7
func Convert(number int) string {

	var collection []int
	var results []string

	// Checking for factors in the given number
	for eachNumber := 1; eachNumber <= number; eachNumber++ {
		if number%eachNumber == 0 {
			collection = append(collection, eachNumber)
		}
	}
	for _, i := range collection {
		if i == 3 {
			results = append(results, "Pling")
		}
		if i == 5 {
			results = append(results, "Plang")
		}
		if i == 7 {
			results = append(results, "Plong")
		}
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
