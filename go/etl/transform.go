package etl

import (
	"strings"
)

// Transform function(map[int][]string) transforms the input map
func Transform(actual map[int][]string) map[string]int {
	expectation := make(map[string]int)

	for key, val := range actual {
		for _, c := range val {
			expectation[strings.ToLower(c)] = key
		}
	}
	return expectation
}
