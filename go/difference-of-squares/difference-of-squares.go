package diffsquares

// Sum(i, where i=1 to n) = n*(n+1)/2
// Sum(i^2, where i=1 to n) = n*(n+1)*(2*n+1)/6,   for n in Natural nums

// SumOfSquares function(int) int
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSums function(int) int
func SquareOfSums(n int) int {
	return (n * (n + 1) / 2) * (n * (n + 1) / 2)
}

// Difference function(int)int
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
