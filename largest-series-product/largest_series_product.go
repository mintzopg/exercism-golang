package lsproduct

import (
	"errors"
	"regexp"
	"strconv"
)

func LargestSeriesProduct(s string, n int) (prod int64, e error) {
	// define an array with possible errors
	// errors {"span must be greater than zero", "span must be smaller than string length", "digits input must only contain digits"}
	p := int64(1) // product per iteration; temporary

	// check span > 0
	if n < 0 {
		return -1, errors.New("span must be greater than zero")
	}
	// check span < len(string)
	if n > len(s) {
		return -1, errors.New("span must be smaller than string length")
	}
	if n == 0 {
		return 1, e
	}

	// check for valid input; contains only digits
	re := regexp.MustCompile(`(?m)\D`)
	if re.MatchString(s) == true {
		return -1, errors.New("digits input must only contain digits")
	}

	for {
		if n > len(s) {
			return prod, e
		}
		for i := 0; i < n; i++ {
			k, _ := strconv.Atoi(string(s[i]))
			p *= int64(k)
		}
		if p > prod {
			prod = p
		}
		p = 1
		s = s[1:]
	}
}
