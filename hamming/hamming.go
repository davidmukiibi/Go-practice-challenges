// Package hamming calculates the hamming sequence of 2 strings when given and
// returns the hamming frequency between the 2 strings
package hamming

import (
	"errors"
)

// Distance function takes in 2 string arguments
// it compares each character to find out if they are equal
// if they are not, it increases count by 1 and eventually
// returns the value of count and/or error if any occurs in the process
func Distance(a, b string) (int, error) {

	// check if strings entered are of the same length
	if len(a) != len(b) {
		err := errors.New("the 2 strings given are not of the same length")
		return 0, err
	}
	// this is where the magic happens
	var count int
	for i := range a {
		if a[i] != b[i] {
			count++
		}
	}
	return count, nil
}
