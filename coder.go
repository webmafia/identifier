package identity

import (
	"strconv"
)

type Coder interface {
	Decode(id string) []uint64
	Encode(numbers []uint64) (string, error)
}

var _ Coder = stringCoder{}

type stringCoder struct{}

// Decode implements Coder.
func (stringCoder) Decode(id string) []uint64 {
	i, _ := strconv.Atoi(id)
	return []uint64{uint64(i)}
}

// Encode implements Coder.
func (stringCoder) Encode(numbers []uint64) (string, error) {
	return strconv.Itoa(int(numbers[0])), nil
}
