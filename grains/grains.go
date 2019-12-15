package grains

import (
	"errors"
)

// Square function(int)(uint64, bool) chessboard grains/square calc
func Square(sq int) (uint64, error) {
	// number of grains in square = 2^n
	if sq < 1 || sq > 64 {
		return 0, errors.New("invalid square number")
	}

	n := sq - 1 // sq >= 1
	out := 1 << uint(n)

	return uint64(out), nil
}

// Total function returns total number of grains
/*
	Sum (a*r^k) = a(1-r^n)/(1-r) ; here a =1, r = 2 hence Sum = 2^n - 1 <==> 1 << n - 1
	k 0 to n-1
*/
func Total() uint64 {
	return 1<<64 - 1
}
