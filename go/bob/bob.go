package bob

import (
	"strings"
	"unicode"
)

const question = "Sure."
const yelling = "Whoa, chill out!"
const yellingQuestion = "Calm down, I know what I'm doing!"
const sayNothing = "Fine. Be that way!"
const anything = "Whatever."

// Hey function (string) string
func Hey(remark string) string {
	rr := []rune(remark)
	l := len(remark)
	var lower, upper, space int

	for _, c := range rr {
		switch {
		case unicode.IsLower(c):
			lower++
		case unicode.IsUpper(c):
			upper++
		case unicode.IsSpace(c):
			space++
		}
	}

	// outputs
	// silence or prolonged silence
	if l == 0 || space == l {
		return sayNothing
	}
	// check for question
	if strings.HasSuffix(strings.TrimRight(remark, " "), "?") {
		// check for yelling question
		if upper > 0 && lower == 0 {
			return yellingQuestion
		} // else is a question
		return question
	}
	// check for yelling
	if upper > 0 && lower == 0 {
		return yelling
	}

	return anything
}
