package wordcount

import (
	"regexp"
	"strings"
)

// Frequency of words appearance in a piece of strings
type Frequency map[string]int

// WordCount (string) Frequency
func WordCount(str string) Frequency {
	freq := make(Frequency, 0)

	// clean the string
	var re = regexp.MustCompile(`['\w]+`)
	ss := re.FindAllString(str, -1) // []string

	for _, word := range ss {
		word = strings.ToLower(word)
		freq[strings.Trim(word, "'")]++
	}

	return freq
}
