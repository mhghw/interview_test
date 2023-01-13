package randomStr

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabetNumber = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func NewRandomString() string {
	b := make([]byte, 100)
	for i := range b {
		b[i] = alphabetNumber[rand.Intn(len(alphabetNumber))]
	}
	return string(b)
}
