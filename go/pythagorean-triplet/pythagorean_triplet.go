package pythagorean

// helper functions
func maxof(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func minof(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

////////////////////////////////////////////////////////

type Triplet [3]int

func Range(min, max int) []Triplet {
	var a, b, c int
	var out []Triplet

	for a = min; a <= max; a++ {
		for b = a; b <= max; b++ {
			for c = b; c <= max; c++ {
				if c*c == a*a+b*b {
					out = append(out, Triplet{minof(a, b), maxof(a, b), c})
				}
			}
		}
	}
	return out
}

func Sum(p int) []Triplet { // this func signature is not as defined in exercism instructions as it would not work
	triplets := Range(1, p)
	out := []Triplet{}

	for _, v := range triplets {
		if v[0]+v[1]+v[2] == p {
			out = append(out, v)
		}
	}
	return out
}
