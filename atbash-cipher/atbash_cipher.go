package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

// Atbash takes a string and returns the Atbash ciphered string
func Atbash(s string) string {
	var b strings.Builder
	// Manipulate input string; (1) regex to only keep letters; (2) to lower case
	reg := regexp.MustCompile(`(?m)[^A-Za-z0-9]`)
	s = strings.ToLower(reg.ReplaceAllLiteralString(s, ""))
	// fmt.Println(s) // comment out

	// string ciphering starts here
	count := 1
	var k rune
	for _, c := range s {
		if unicode.IsNumber(c) {
			k = 2 * c // unicode values
		} else {
			k = 219 // 122+97 unicode values
		}
		if count > 5 {
			b.WriteRune(' ')
			b.WriteRune(k - c)
			count = 2
		} else {
			b.WriteRune(k - c)
			count++
		}
	}
	return b.String()
}
