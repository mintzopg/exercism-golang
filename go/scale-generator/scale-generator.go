package scale

import (
	"strings"
)

// Scale function(string, string)[] string
func Scale(tonic, interval string) []string {
	scale := []string{}
	out := []string{}

	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}

	for i := range scale {
		if strings.Title(tonic) == scale[i] {
			// no interval given
			if len(interval) == 0 { // return 12-note chromic scale
				out = append(scale[i:], scale[:i]...)
				break
			} else {
				// interval is given
				out = append(out, scale[i]) // add the tonic pitch and continue to generate the scale
				// take the intervals and add notes accordingly
				for j := 0; j < len(interval)-1; j++ {
					switch string(interval[j]) {
					case "M":
						scale = append(scale, scale[:i]...) // wrap the scale slice
						// whole step
						i += 2
						out = append(out, scale[i])
					case "m":
						scale = append(scale, scale[:i]...) // wrap the scale slice
						// half step
						i++
						out = append(out, scale[i])
					case "A":
						scale = append(scale, scale[:i]...) // wrap the scale slice
						// augmented first
						i += 3
						out = append(out, scale[i])
					}

				}
			}
		}
	}
	return out
}
