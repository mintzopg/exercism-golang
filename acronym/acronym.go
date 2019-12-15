package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate function (string)string
func Abbreviate(s string) string {
	s = strings.TrimSpace(s)
	chars := []rune(s)
	out := ""

	if unicode.IsLetter(chars[0]) {
		out += string(unicode.ToUpper(chars[0]))
	}

	for j := 1; j < len(chars); {
		if unicode.IsSpace(chars[j]) || (string(chars[j]) == "-") {
			out += string(unicode.ToUpper(chars[j+1]))
			j += 2
		} else {
			j++
		}
	}

	return out
}
