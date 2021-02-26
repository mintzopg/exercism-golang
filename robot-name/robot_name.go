package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

var names = map[string]bool{} // map to hold active robot names

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Robot structure describes Robot
type Robot struct {
	name string
}

// Name generates a valid name for Robot
func (r *Robot) Name() (string, error) {
	if r.name != "" { // this comes from use of Name() method in test files
		return r.name, nil
	}
	for { // loop will go on until it gets a new valid name
		r.name = newName()               // (1) get a valid random name
		if _, ok := names[r.name]; !ok { // (2) check if already given
			names[r.name] = true // (3) if not add it to cache; the value true or false doesn't matter
			return r.name, nil   // (4) return the string and nil error
		}
	}
}

// Reset simply deletes robot name
func (r *Robot) Reset() {
	r.name = ""
}

func newName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
