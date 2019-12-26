package clock

import (
	"fmt"
)

// Clock API:
//
// type Clock                      // define the clock type
// New(hour, minute int) Clock     // a "constructor"
// (Clock) String() string         // a "stringer"
// (Clock) Add(minutes int) Clock
// (Clock) Subtract(minutes int) Clock
//
// To satisfy the README requirement about clocks being equal, values of
// your Clock type need to work with the == operator. This means that if your
// New function returns a pointer rather than a value, your clocks will
// probably not work with ==.
//
// While the time.Time type in the standard library (https://golang.org/pkg/time/#Time)
// doesn't necessarily need to be used as a basis for your Clock type, it might
// help to look at how constructors there (Date and Now) return values rather
// than pointers. Note also how most time.Time methods have value receivers
// rather than pointer receivers.

// Clock type
type Clock struct {
	h, m int
}

// New Clock constructor method for Clock type
func New(hour, minute int) Clock {
	var h, m int
	h = (hour + int(minute/60)) % 24

	if h < 0 {
		h += 24
	}
	m = minute % 60
	if m < 0 {
		m += 60
		h--
		if h < 0 {
			h += 24
		}
	}
	return Clock{h, m}
}

// Add (minutes) method for Clock type
func (c Clock) Add(minutes int) Clock {
	return New(c.h, c.m+minutes)
}

// Subtract (minutes) method for Clock type
func (c Clock) Subtract(minutes int) Clock {
	return New(c.h, c.m-minutes)
}

// toString() method for Clock type
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
