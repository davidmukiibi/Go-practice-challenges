package acronym

import (
	s "strings"
)

// Abbreviate takes in a string of words and returns the first letter on every word
func Abbreviate(word string) string {

	var storage []string
	var acronym string

	// checks if the phrase contains "-"
	if s.Contains(word, "-") {
		word = s.Replace(word, "-", " ", -1)
		words := s.SplitAfterN(word, " ", -1)
		for i := 0; i < len(words); i++ {
			w := s.SplitAfterN(words[i], "", -1)
			c := s.ToUpper(w[0])
			storage = append(storage, c)
		}
		acronym = s.Join(storage, "")
		// checks if the acronym contains "-"
		if s.Contains(acronym, " ") {
			acronym = s.Replace(acronym, " ", "", -1)
		}
	} else {
		// splits sentence into separate words
		words := s.SplitAfterN(word, " ", -1)
		for i := 0; i < len(words); i++ {
			w := s.SplitAfterN(words[i], "", -1)
			c := s.ToUpper(w[0])
			storage = append(storage, c)
		}
		acronym = s.Join(storage, "")
	}

	return acronym

}
