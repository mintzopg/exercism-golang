// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	Equ Kind = iota // equilateral
	Iso             // isosceles
	Sca             // scalene
	NaT             // not a triangle
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) ||
		math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) ||
		math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return NaT
	case a <= 0 || b <= 0 || c <= 0:
		return NaT
	case a+b < c || a+c < b || b+c < a:
		return NaT
	case a == b && b == c:
		return Equ
	case a == b || a == c || b == c:
		return Iso
	}
	return Sca
}
