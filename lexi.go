package lexi

import (
	"encoding/hex"
	"fmt"
	"math"
)

func PackHex(n uint) string {
	buf := Pack(n)
	return hex.EncodeToString(buf)
}

func Pack(n uint) []byte {

	var bytes []byte
	max := byte(251)
	x := n - uint(max)

	if n < uint(max) {
		bytes = []byte{byte(n)}
	} else if x < 256 {
		bytes = []byte{max, byte(x)}
	} else if x < 256*256 {
		bytes = []byte{max + 1, byte(x / 256), byte(x % 256)}
	} else if x < 256*256*256 {
		bytes = []byte{
			max + 2,
			byte(x / 256 / 256),
			byte((x / 256) % 256),
			byte(x % 256),
		}
	} else if x < 256*256*256*256 {
		bytes = []byte{
			max + 3,
			byte(x / 256 / 256 / 256),
			byte((x / 256 / 256) % 256),
			byte((x / 256) % 256),
			byte(x % 256),
		}
	} else {
		exp := (math.Log(float64(x)) / math.Log(2)) - 32
		bytes = []byte{255}
		bs := Pack(uint(exp))
		bytes = append(bytes, bs...)
		res := float64(x) / math.Pow(2, exp-11)
		bytes = append(bytes, bytesOf(x/uint(res))...)
	}

	return bytes
}

func bytesOf(x uint) []byte {
	bytes := []byte{}
	d := uint(1)
	for i := 0; i < 6; i++ {
		bytes = append([]byte{byte((x / d) % 256)}, bytes...)
	}
	return bytes
}

func UnpackHex(str string) (uint, error) {
	buf, err := hex.DecodeString(str)
	if err != nil {
		return 0, err
	}

	return Unpack(buf)
}

func Unpack(xs []byte) (uint, error) {

	length := len(xs)

	if length == 0 {
		return 0, nil
	} else if length == 1 && xs[0] < 251 {
		return uint(xs[0]), nil
	} else if length == 2 && xs[0] == 251 {
		return 251 + uint(xs[1]), nil
	} else if length == 3 && xs[0] == 252 {
		return 251 + 256*uint(xs[1]) + uint(xs[2]), nil
	} else if length == 4 && xs[0] == 253 {
		return 251 + 256*256*uint(xs[1]) + 256*uint(xs[2]) + uint(xs[3]), nil
	} else if length == 5 && xs[0] == 254 {
		return 251 + 256*256*256*uint(xs[1]) + 256*256*uint(xs[2]) + 256*uint(xs[3]) + uint(xs[4]), nil
	} else if length > 5 && xs[0] == 255 {
		m := uint(0)
		x := uint(1)
		pivot := max(2, uint(length-6))
		for i := uint(length) - 1; i >= pivot; i-- {
			m += x * uint(xs[i])
			x *= 256
		}
		var n uint
		var err error
		if xs[1]+32 < 251 {
			n, err = Unpack([]byte{xs[1], xs[2] + 21})
			if err != nil {
				return 0, err
			}
			n -= 11
		} else if xs[0] == 255 && xs[1] < 251 {
			n = uint(xs[1]) + 21
		} else if pivot == 3 {
			n, err = Unpack([]byte{xs[1], xs[2] + 21})
			if err != nil {
				return 0, err
			}
		} else if pivot == 4 {
			n, err = Unpack([]byte{xs[1], xs[2], xs[3] + 21})
			if err != nil {
				return 0, err
			}
		}
		return 251 + m/uint(math.Pow(2, float64(32-n))), nil
	}

	return 0, fmt.Errorf("Not valid lexi string")
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
