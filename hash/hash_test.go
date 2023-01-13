package hash_test

import (
	"int_test/hash"
	"testing"
)

type TestCase struct {
	str  string
	hash string
}

var TestCases = []TestCase{
	{"tAEOo8cvP08NjqK9", "d28c243ddc3ad3c126366fb5fe805044a995adb2ba0e62a6cd0467ee650f3d5f"},
	{"fRccKErSixYXZmif", "47f1b206f452699e39812384706450983d97bbc5f922f870958182e72fe6c8a3"},
	{"2U1VP1nbpE0LR4Pi", "88a0572cf85e7811264acb63a01de3653bdbad793d8ed2472d1ea78b2502f1d6"},
	{"bB5wRCBAUdw0LSRh", "a2592da3ac2878fda952a1a0f25cf5c78d28b9b634e9c748c7d7836ce02b0009"},
	{"V0MyIhfCCBGMirHV", "8bbb30a81da594a0fe36b31cb3858c2db70f7ab9bb2fb961ccc1fb93a11ce4bb"},
	{"DClpPkdjKMw5w2DK", "90a9c0f6160bf976b4528f02b79f7c963243395f9949fe76f6f09f738aade421"},
	{"L7EMTS7OHnLEo11y", "74330b4b7c32cb46ddde4e2a0840ed55cf3313b06f17ae2d9a4aec4c220b6c4c"},
}

func TestHashString(t *testing.T) {
	t.Parallel()
	for _, testCase := range TestCases {
		if h := hash.HashString(testCase.str); testCase.hash != h {
			t.Errorf("for string: %v got hash: %v but expected: %v", testCase.str, h, testCase.hash)
		}
	}
}
