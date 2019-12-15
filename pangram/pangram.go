package pangram

import (
	"strings"
)

var letters = map[rune]bool{}

// IsPangram function(string)bool
func IsPangram(s string) bool {
	s = strings.ToLower(s) // case insensitive
	// initialize map
	for i := 'a'; i <= 'z'; i++ {
		letters[i] = false
	}

	for _, b := range s {
		// if letter flag as found
		_, ok := letters[b]
		if ok {
			letters[b] = true
		}
	}

	flag := true
	for _, val := range letters {
		flag = flag && val
	}

	return flag
}
