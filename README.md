# lexi-int

Lexicographical ordering for integers. Produces `[]byte` and hex `string`
encodings of integers that maintain the same ordering before and after encoding.

From the implementation by
[substack](https://github.com/substack/lexicographic-integer).

## example

```go
import (
    lexi "github.com/filwisher/lexi-int"
)

func property(a, b uint32) bool {

    astr := lexi.PackHex(uint(a))
    bstr := lexi.PackHex(uint(b))

    if a < b {
        return astr < bstr
    } else if a > b {
        return astr > bstr
    } else {
        return astr == bstr
    }
}
```

## methods

### lexi.Pack(uint) []byte
### lexi.Unpack([]byte) (uint, error)
### lexi.PackHex(uint) string
### lexi.UnpackHex(string) (uint, error)

## install
```
$ go get https://github.com/substack/lexicographic-integer
```
