package randomStr

import (
	"log"
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
	log.Println(string(b))
	return string(b)
}
