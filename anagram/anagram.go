package anagram

import "strings"

// perm recursively generates all permutations of a []rune and applies function f on them
func perm(a []rune, f func([]rune), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Detect anagrams of the input string subject
func Detect(subject string, candidates []string) []string {
	out := make([]string, 0)                            // slice of candidates to return
	var m = map[string]int{strings.ToLower(subject): 0} // helper map to eliminate duplicates

	f := func(a []rune) { // check each permutation against each candidate string; case-insensitive comparison
		for _, s := range candidates {
			if strings.ToLower((string(a))) == strings.ToLower(s) {
				_, ok := m[strings.ToLower(s)]
				if !ok {
					m[strings.ToLower(s)]++
					out = append(out, s)
				}
			}
		}
	}

	perm([]rune(subject), f, 0)
	return out
}
