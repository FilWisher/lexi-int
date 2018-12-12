package lexi

import (
	"testing"
	"testing/quick"
)

func TestPackHex(t *testing.T) {
	f := func(a, b uint32) bool {
		astr := PackHex(uint(a))
		bstr := PackHex(uint(b))

		if a < b {
			return astr < bstr
		} else if a > b {
			return astr > bstr
		} else {
			return astr == bstr
		}
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

func TestInverse(t *testing.T) {
	f := func(a uint32) bool {
		str := PackHex(uint(a))

		n, err := UnpackHex(str)
		if err != nil {
			t.Log(err)
			return false
		}

		return n == uint(a)
	}

	if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}
