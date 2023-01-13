package hash

import (
	"crypto/sha256"
	"fmt"
)

var hasher = sha256.New()

func HashString(s string) string {
	hasher.Write([]byte(s))
	hashed := hasher.Sum(nil)
	hasher.Reset()
	return fmt.Sprintf("%x", hashed)
}
