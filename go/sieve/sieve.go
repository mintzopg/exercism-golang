package sieve

// Sieve function Sieve(n int) []int
func Sieve(n int) []int {
	var out []int
	prime := make([]bool, n+1)
	// initilaize array of booleans
	for i := 0; i < n+1; i++ {
		prime[i] = true
	}
	p := 2
	for p*p <= n {
		if prime[p] == true {
			// update all multiples of p
			for i := p * 2; i < n+1; i += p {
				prime[i] = false
			}
		}
		p++
	}
	// get all prime numbers
	for p := 2; p <= n; p++ {
		if prime[p] {
			out = append(out, p)
		}
	}
	return out
}
