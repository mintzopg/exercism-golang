package luhn

import (
	"strings"
	"unicode"
)

func isValid(r rune) bool {
	if unicode.IsDigit(r) || unicode.IsSpace(r) {
		return true
	}
	return false
}

// Valid function(string)bool
func Valid(s string) bool {
	ss := []rune(strings.TrimSpace(s))
	if len(ss) <= 1 {
		return false
	}

	// work with valid digits
	v := make([]int, 0, cap(ss)) // v empty slice of ints with same cap as ss
	for _, x := range ss {
		if !isValid(x) {
			return false
		} else if isValid(x) && (!unicode.IsSpace(x)) {
			v = append(v, int(x)-'0') // append int; note: all digits will be in range
		}
	}

	// implement the algo
	// step (1) double every second digit
	for i := len(v) - 2; i >= 0; i -= 2 {
		if v[i]*2 > 9 {
			v[i] = v[i]*2 - 9
		} else {
			v[i] *= 2
		}
	}

	// step (2) summ all digits in v
	sum := 0
	for _, x := range v {
		sum += x
	}

	// step (3) if sum % 10 == 0 -> valid
	if !(sum%10 == 0) {
		return false
	}

	return true
}
