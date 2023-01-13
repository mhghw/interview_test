package randomStr_test

import (
	randomStr "int_test/random"
	"regexp"
	"testing"
)

var validRegexp = `^[A-Za-z0-9]*$`

func TestNewRandomString(t *testing.T) {
	t.Parallel()
	validString := regexp.MustCompile(validRegexp)

	for i := 0; i < 200; i++ {
		str := randomStr.NewRandomString()
		if !validString.MatchString(str) {
			t.Errorf("invalid string generated: %v", str)
		}
	}
}
