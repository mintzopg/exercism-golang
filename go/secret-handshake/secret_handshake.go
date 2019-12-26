package secret

import (
	"fmt"
	"strings"
)

// helper func sliceReverse
func sliceReverse(a []string) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

// Handshake func(uint) []string, returns secret codes
func Handshake(code uint) []string {
	out := []string{}
	codeString := fmt.Sprintf("%032b", code)
	codeString = strings.TrimLeftFunc(codeString, func(r rune) bool {
		return r == '0'
	})
	codes := [4] string{"wink", "double blink", "close your eyes", "jump"}

	N := len(codeString)
	for i := 1; i <= 4; i++{
		if N-i >= 0 && codeString[N-i] == '1' {
			out = append(out, codes[i-1])
		}
	}
	if N-5 >= 0 && codeString[N-5] == '1' {
		sliceReverse(out)
	}
	
	return out
}
