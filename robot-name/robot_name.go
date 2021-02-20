package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// for numbers
var minNum = 100
var maxNum = 1000

// for ascii capital letters
var minASCII = 65 // (A=65)
var maxASCII = 91 // (Z=90)

var names = map[string]bool{} // map to hold active robot names

func init() {
	rand.Seed(time.Now().UnixNano())
}

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
		r.name = createValidName()       // (1) get a valid random name
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

/////////////////////////////////////////////////////////////////////////////////
func randomInt(min, max int) int { // generate random number in range [100, 1000)
	return rand.Intn(max-min) + min
}

func createValidName() (name string) {
	bytes := make([]byte, 2)
	num := randomInt(minNum, maxNum) // get number in [100, 999]
	for i := 0; i < 2; i++ {         // get two capital letters
		bytes[i] = byte(randomInt(minASCII, maxASCII))
	}
	return string(bytes) + fmt.Sprintf("%d", num)
}
