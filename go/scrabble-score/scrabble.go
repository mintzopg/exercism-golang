package scrabble

import "strings"

// Score function (string) int   returns scrabble score
func Score(str string) int {
	// define map for the scores
	score := 0
	m := map[string]int{
		"A, E, I, O, U, L, N, R, S, T": 1,
		"D, G":                         2,
		"B, C, M, P":                   3,
		"F, H, V, W, Y":                4,
		"K":                            5,
		"J, X":                         8,
		"Q, Z":                         10,
	}

	// for every letter in input string
	chars := strings.Split(strings.ToUpper(str), "")
	for _, ch := range chars {
		// look up in every key of the map if it is contained
		for key, val := range m {
			// if yes add the score
			if strings.Contains(key, ch) {
				score += val
			}
		}
	}
	return score
}
