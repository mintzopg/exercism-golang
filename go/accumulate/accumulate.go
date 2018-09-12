package accumulate

// Accumulate function([]string, func)[]string
func Accumulate(a []string, f func(string) string) []string {
	b := make([]string, len(a))

	for i, x := range a {
		b[i] = f(x)
	}

	return b
}
