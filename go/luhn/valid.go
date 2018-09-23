package luhn

import "unicode"

// Valid function(string)bool
func Valid(s string) bool {
	ss := []rune(s)
	if len(ss) <= 1 {
		return false
	}

	// any rune is invalid if not digit or space
	isValid := func(r rune) bool {
		if unicode.IsDigit(r) || unicode.IsSpace(r) {
			return true
		}
		return false
	}

	// work with valid digits
	v := ss[:0] // v empty slice with same cap as ss
	for _, x := range ss {
		if isValid(x) && (!unicode.IsSpace(x)) {
			v = append(v, x)
		} else {
			return false
		}
	}

	return true
}
