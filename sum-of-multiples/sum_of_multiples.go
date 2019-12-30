package summultiples

import "sort"

// SumMultiples returns the sum of the multiples of the given divisors up to, but not including, the given limit.
//func SumMultiples(int, ...int)int

// SumMultiples (N int, divisors ...int) int // return the sum of multiples of divisors < N
func SumMultiples(limit int, divisors ...int) int {
	if len(divisors) == 0 { // no factors means empty sum
		return 0
	}
	if len(divisors) == 1 && divisors[0] == 0 {
		return 0 // the only multiple of 0 is 0
	}

	var sum int
	vec := []int{}

	for _, d := range divisors { // range over factors
		if d == 0 { // 0 factor does not affect the sum, skip
			break
		}
		multiples := []int{}
		for j := d; j < limit; j++ { // check for multiples
			if j%d == 0 {
				multiples = append(multiples, j)
			}
		}
		vec = append(vec, multiples...)
	}

	// deduplicate in place
	if len(vec) > 1 {
		sort.Ints(vec)
		j := 0
		for i := 1; i < len(vec); i++ {
			if vec[j] == vec[i] {
				continue
			}
			j++
			vec[j] = vec[i]
		}
		vec = vec[:j+1]
	}

	// compute the sum of multiples
	for _, k := range vec {
		sum += k
	}
	return sum
}
