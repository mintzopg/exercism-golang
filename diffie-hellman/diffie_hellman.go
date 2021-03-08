package diffiehellman

import (
	crand "crypto/rand"
	"encoding/binary"
	"math/big"
	mrand "math/rand"
)

// math/rand Source is an interface. Need to implement rand64bitSrc to return a 64 bit
// random source for use in big.Rand
/* type Source interface {
     Int63() int64
     Seed(seed int64)
}
*/

type rand64bitSrc struct{} // implements 64 bit random source

func (s rand64bitSrc) Seed(seed int64) {
	// don't need Seed in crypto/rand
}

// Uint64 returns a 64 bit random number
// using encoding/binary to read 8 bytes from io.Reader in crypto/rand
func (s rand64bitSrc) Uint64() (value uint64) {
	binary.Read(crand.Reader, binary.BigEndian, &value)
	return value
}

func (s rand64bitSrc) Int63() int64 {
	return int64(s.Uint64() & ^uint64(63))
}

//PrivateKey generates a random private key (step 1)
func PrivateKey(p *big.Int) *big.Int {
	// where 1 < key < p
	key := new(big.Int)
	uplimit := new(big.Int).Sub(p, big.NewInt(2))
	rs := mrand.New(new(rand64bitSrc))
	return key.Rand(rs, uplimit).Add(key, big.NewInt(2))
}

// PublicKey (step 2)
func PublicKey(private, p *big.Int, g int64) *big.Int {
	// A = g**a mod p
	key := &big.Int{}
	return key.Exp(big.NewInt(g), private, p)
}

// NewPair (step 2.1)
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey (step 3)
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
