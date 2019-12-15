package raindrops

import "strconv"

// Convert function(int) string
func Convert(num int) string {
	var outstring string

	if num%3 == 0 {
		outstring += "Pling"
	}
	if num%5 == 0 {
		outstring += "Plang"
	}
	if num%7 == 0 {
		outstring += "Plong"
	}

	if len(outstring) == 0 {
		outstring = strconv.Itoa(num)
		return outstring
	}
	return outstring
}
