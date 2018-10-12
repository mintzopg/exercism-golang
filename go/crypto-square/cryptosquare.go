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

// find cryptosquare dimensions
func findDim(n int) (cols int, rows int) {
	// n number of chars; copy to float64(k) required
	k := float64(n)
	var r float64
	c := math.Sqrt(k)
	k++
	for {
		if c == math.Trunc(c) { // if c is integer return it
			break
		} else {
			k++
			c = math.Sqrt(k)
		}
	}
	if c*c == float64(n) {
		r = c
	} else {
		r = c - 1
	}
	return int(c), int(r)
}

// Encode function(normalized string)string
func Encode(txt string) string {
	// normalize the input
	txt = Normalize(txt)
	n := len(txt)
	stringOut := ""
	
	// check the special cases
	if n == 0 {
		return ""
	}
	if n == 1 {
		return txt
	}

	c, r := findDim(n)
	// pad with spaces at the end if necessary
	if n != c*r {
		txt += strings.Repeat(" ", c*r-n)
	}

	square := make([][]string, c)
	for j := 0; j < r; j++ {
		for i := 0; i < c; i++ {
			square[i] = append(square[i], string(txt[i]))
		}
		txt = txt[c:]
	}
	// construct output string
	for i := range square {
		stringOut += strings.Join(square[i], "") + " "
	}
	
	return stringOut[:len(stringOut)-1]
}
