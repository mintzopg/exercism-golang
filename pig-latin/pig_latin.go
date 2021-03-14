package piglatin

import (
	"regexp"
	"strings"
)

// not worrying about capitals since no capitals in test cases
// assuming that sentences will space separated words

// Sentence (string) string, translates a string from English to Pig latin
func Sentence(s string) string {
	var b strings.Builder          // create a builder to compose the output string
	words := strings.Split(s, " ") // split input strings to get the words

	// define regexp for each rule
	re1 := regexp.MustCompile(`(?m)^[aeiou]|^(xr)|^(yt)`)          // Rule #1; Vowel sounds
	re3 := regexp.MustCompile(`(?m)^[bcdfghjklmnpqrstvwxz]*(qu)`)  // Rule #3; capture this first
	re2 := regexp.MustCompile(`(?m)^[bcdfghjklmnpqrstvwxz]+|^[y]`) // Rule #2; y is treated like a consonant at the beginning oof a word
	// re4 := regexp.MustCompile(`(?m)^([bcdfghjklmnpqrstvwxz]+y)|^.$`) // Rule #4; covered by previous rules

	var sp string // separator space to set in case of phrase len(words) > 1
	for i, word := range words {
		if i < len(words)-1 {
			sp = " "
		} else {
			sp = ""
		}
		switch {
		case re1.MatchString(word):
			b.WriteString(word + "ay" + sp)
		case re3.MatchString(word):
			b.WriteString(re3.ReplaceAllString(word, "") + re3.FindString(word) + "ay" + sp)
		case re2.MatchString(word):
			b.WriteString(re2.ReplaceAllString(word, "") + re2.FindString(word) + "ay" + sp)
			// case re4.MatchString(word):
			// 	b.WriteString(re4.ReplaceAllString(word, "") + re4.FindString(word) + "ay" + sp)
		}
	}
	return b.String()
}
