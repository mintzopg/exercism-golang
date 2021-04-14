package twelve

import (
	"fmt"
	"strings"
)

var days = []string{
	"first", "second", "third", "fourth", "fifth", "sixth",
	"seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
}

var gifts = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

<<<<<<< HEAD
// Song generates the entire song
=======
>>>>>>> master
func Song() string {
	var sb strings.Builder
	for i := 1; i <= 11; i++ {
		sb.WriteString(Verse(i) + "\n")
	}
	sb.WriteString(Verse(12)) // write without new line
	return sb.String()
}

<<<<<<< HEAD
// Verse takes a day number and produces the corresponding verse
=======
>>>>>>> master
func Verse(i int) string {
	var verse string
	if i < 1 || i > len(days) {
		panic("day should be between 1 and 12")
	}
	k := i - 1 // for convenience
	if k == 0 {
		verse = fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[k], gifts[k])
	} else {
		v := []string{}
		for j := k; j > 0; j-- {
			v = append(v, gifts[j])
		}
		v = append(v, fmt.Sprintf("and %s", gifts[0]))
		verse = fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[k], strings.Join(v, ", "))
	}
	return verse
}
