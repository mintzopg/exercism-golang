package tournament

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// map to hold league data => "team name" [MP W D L P]
var mp = make(map[string]*[5]int)

// Tuple is used to sort the table of results
type Tuple struct {
	Team string
	MP   int
	W    int
	D    int
	L    int
	P    int
}

type PairList []Tuple

// implement data interface; https://golang.org/pkg/sort/
func (p PairList) Len() int      { return len(p) }
func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool {
	if p[i].P == p[j].P { // on equal score sort alphabetically
		return p[i].Team < p[j].Team
	}
	return p[i].P > p[j].P // sort on score descending
}

// Tally the results of a small football competition
func Tally(r io.Reader, w io.Writer) error {
	// it's a new league it's a new day
	mp = make(map[string]*[5]int)
	// take input into a []string
	var input []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") { // ignore new/empty lines and comments
			continue
		}
		input = append(input, scanner.Text())
	}
	// work on input => validate & count scores
	re := regexp.MustCompile(`(?m)(^[A-Za-z\s]+(;)[A-Za-z\s]+(;))((win)|(loss)|(draw))$`)
	for _, line := range input {
		// check if is valid entry; return error if not
		if !re.MatchString(line) {
			return fmt.Errorf("invalid input for entry %q", line)
		}

		l := strings.SplitAfterN(line, ";", 3) // []string => [team1; team2; win|loss|draw]
		team1 := l[0][:len(l[0])-1]
		team2 := l[1][:len(l[1])-1]
		outcome := l[2]

		switch outcome {
		case "win": // team1 wins - team2 losses
			updateData(team1, "win")
			updateData(team2, "loss")
		case "loss": // team1 losses - team2 wins
			updateData(team1, "loss")
			updateData(team2, "win")
		case "draw": // it's a draw
			updateData(team1, "draw")
			updateData(team2, "draw")
		}
	}
	// format output
	p := make(PairList, len(mp))
	i := 0
	for k, v := range mp {
		p[i] = Tuple{k, v[0], v[1], v[2], v[3], v[4]}
		i++
	}
	sort.Sort(p) //p is sorted by score

	fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", "Team", "MP", " W", "D", "L", "P")
	for _, k := range p {
		fmt.Fprintf(w, "%-31v|%3v |%3v |%3v |%3v |%3v\n", k.Team, k.MP, k.W, k.D, k.L, k.P)
	}
	return nil
}

func updateData(team, outcome string) {
	switch outcome {
	case "win":
		if v, ok := mp[team]; !ok { // first time entry
			mp[team] = &[5]int{1, 1, 0, 0, 3}
		} else {
			v[0]++
			v[1]++
			v[4] += 3
		}
	case "loss":
		if v, ok := mp[team]; !ok { // first time entry
			mp[team] = &[5]int{1, 0, 0, 1, 0}
		} else {
			v[0]++
			v[3]++
		}
	case "draw":
		if v, ok := mp[team]; !ok { // first time entry
			mp[team] = &[5]int{1, 0, 1, 0, 1}
		} else {
			v[0]++
			v[2]++
			v[4] += 1
		}
	}
}
