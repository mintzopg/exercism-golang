package encode

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode encodes the given string
func RunLengthEncode(s string) string {
	if len(s) == 0 {
		return ""
	}
	count := make([]int, len(s))  // count occurences
	chars := make([]rune, len(s)) // characters
	i := 0                        // hold index in count, chars slices

	// initialize & encode
	pre := []rune(s)[0]
	chars[i] = pre
	count[i]++

	for _, cur := range []rune(s)[1:] { // range over the rest of chars
		if cur != pre { // new char
			i++ // move to next position in slices
			chars[i] = cur
			count[i]++
			pre = cur
		} else {
			count[i]++ // if same as previous add counter
		}
	}

	// construct encoded string
	var b strings.Builder
	for j := 0; j < len(s); j++ {
		if count[j] == 0 {
			break
		} else if count[j] == 1 {
			fmt.Fprintf(&b, "%v", string(chars[j]))
		} else {
			fmt.Fprintf(&b, "%d%v", count[j], string(chars[j]))
		}
	}

	return b.String()
}

// RunLengthDecode decodes the given string
func RunLengthDecode(s string) string {
	var b strings.Builder

	findNum := func(c rune) bool { // detect number
		return unicode.IsNumber(c)
	}
	findLet := func(c rune) bool { // detect letter or space
		return unicode.IsLetter(c) || unicode.IsSpace(c)
	}

	for len(s) > 0 { // s to be mutated inside the loop
		n := strings.IndexFunc(s, findNum)
		l := strings.IndexFunc(s, findLet)
		// either l > n or l < n
		if l > n { // number precedes letter
			num, _ := strconv.Atoi(s[:l]) // get number from string
			if num == 0 {
				num = 1 // if num = 0, print the character one time
			}
			fmt.Fprintf(&b, "%v", strings.Repeat(string(s[l]), num))
			if l+1 <= len(s) {
				s = s[l+1:]
			} else {
				break
			}
		} else if l < n { // letter precedes number
			fmt.Fprintf(&b, "%v", string(s[l]))
			if l+1 <= len(s) {
				s = s[l+1:]
			} else {
				break
			}
		} else { // l = n = -1 means ""
			fmt.Fprintf(&b, "%v", "")
		}
	}
	return b.String()
}
