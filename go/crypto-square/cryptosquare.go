package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Normalize function(string)string; returns only letters in lowe case & no spaces
func Normalize(txt string) string {
	chars := []rune(txt)
	var sb strings.Builder

	for _, c := range chars {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			sb.WriteRune(unicode.ToLower(c))
		}
	}
	return sb.String()
}

func findCol(n int) int {
	// n number of chars
	c := math.Sqrt(float64(n))
	n++
	for {
		if c == math.Trunc(c) { // if c is integer return it
			break
		} else {
			n++
			c = math.Sqrt(float64(n))
		}
	}
	return int(c)
}

// Encode function(normalized string)string
func Encode(txt string) string {
	// normalize the input
	txt = Normalize(txt)
	n := len(txt)

	// check the special cases
	if n == 0 {
		return ""
	}
	if n == 1 {
		return txt
	}

	// finc number of rows, columns
	r := 0
	c := findCol(n)
	if c*c == n {
		r = c
	} else {
		r = c - 1
	}

	// pad with spaces at the end if necessary
	if n != c*r {
		txt += strings.Repeat(" ", n-c*r)
	}

	square := make([]string, c)

}
