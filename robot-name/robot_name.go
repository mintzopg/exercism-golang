package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const uniqueNames = 26 * 26 * 10 * 10 * 10

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
	if len(names) == uniqueNames {
		return "", errors.New("the number of possible unique names is exhausted")
	}
	r.name = newName()  // (1) get a valid random name
	for names[r.name] { // (2) name in cache? true
		r.name = newName() // (3) repeat with new name
	}
	names[r.name] = true // add in cache
	return r.name, nil
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
