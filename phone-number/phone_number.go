package phonenumber

import (
	"errors"
	"regexp"
	"strings"
	"unicode"
)

// Number takes in a string and returns a telephone number
func Number(s string) (string, error) {
	// get a pure string containing only digits
	f := func(c rune) bool {
		return !unicode.IsDigit(c)
	}
	numstr := strings.Join(strings.FieldsFunc(s, f), "") // string of digits

	// get rid of country code if present
	if strings.HasPrefix(numstr, "1") {
		numstr = numstr[1:]
	}

	// validate the number against the valid format NXXNXXXXXX,
	// where N ε [2-9] and X ε [0-9]
	var re = regexp.MustCompile(`(?m)^(1|)[2-9]\d{2}[2-9]\d{6}$`)
	if re.MatchString(numstr) {
		return numstr, nil
	}
	return "", errors.New("invalid phone number" + s)
}

// AreaCode returns the area code
func AreaCode(s string) (string, error) {
	num, err := Number(s)
	if err == nil {
		return num[:3], nil
	}
	return "", err
}

// Format formats a telephone number
func Format(s string) (string, error) {
	num, err := Number(s) // ok == false means no error
	if err == nil {
		areacode, _ := AreaCode(s)
		return "(" + areacode + ") " + num[3:6] + "-" + num[6:], nil
	}
	return "", err
}
