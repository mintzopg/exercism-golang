package scrabble

import "unicode"

var m = map[rune]int{
	'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
	'D': 2, 'G': 2, 'B': 3, 'C': 3, 'M': 3, 'P': 3, 'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
	'K': 5, 'J': 8, 'X': 8, 'Q': 10, 'Z': 10,
}

// Score function (string) int   returns scrabble score
func Score(str string) int {
	// define map for the scores
	score := 0

	// for every rune in input string
	for _, ch := range str {
		// look up if it is a key of the map and if yes add to score accordingly
		if _, ok := m[unicode.ToUpper(ch)]; ok == true {
			score += m[unicode.ToUpper(ch)]
		}
	}
	return score
}
