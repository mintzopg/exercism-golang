package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram function(string) string // check whever input string is an isogram
func IsIsogram(s string) bool {
	rr := []rune(strings.ToLower(s)) // break string into slice of lower case runes

	var m = map[rune]int{}

	for _, c := range rr {
		if _, ok := m[c]; ok == true && unicode.IsLetter(c) {
			// if rune is a letter AND rune is already in the map -> increase its value
			m[c]++
		} else if unicode.IsLetter(c) {
			// else if rune is a letter -> put it in the map and start the count
			m[c] = 1
		}
	}
	// now map should contain only letters
	// run through the map and if a value is > 1 it is not isogram
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
