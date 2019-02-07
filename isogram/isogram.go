package isogram

import (
	"strings"
)

// IsIsogram function takes i a word and checks it if it has any duplicate letters in in order
// to determine if the word is an Isogram or not.
func IsIsogram(word string) bool {
	var newMap = map[rune]bool{}

	word = strings.ToLower(word)
	for _, i := range word {
		if i == ' ' || i == '-' {
			continue
		}
		if newMap[i] {
			return false
		}
		newMap[i] = true
	}
	return true
}
