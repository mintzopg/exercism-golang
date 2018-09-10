// Package gigasecond provides a function that takes in a time.Time value and returns the time plus 10^9
package gigasecond

// import path for the time package from the standard library
import "time"

const gigasecond = time.Second * 1e9

// AddGigasecond function(t time.Time) time.Time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
