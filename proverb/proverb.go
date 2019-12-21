package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	out := []string{}
	n := len(rhyme)

	if n == 0 {
		return out
	}

	i := 0
	for {
		if i == n-1 {
			s := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
			out = append(out, s)
			return out
		}
		s := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		out = append(out, s)
		i++
	}
}
