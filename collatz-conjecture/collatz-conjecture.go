package collatzconjecture

import "errors"

// CollatzConjecture function(int)(int, error)
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("true")
	}

	k := 0
	var conj func(int)

	conj = func(n int) {
		if n == 1 {
			return
		} else if n%2 == 0 {
			n /= 2
			k++
		} else {
			n = n*3 + 1
			k++
		}
		conj(n)
	}
	conj(n)
	return k, nil
}
