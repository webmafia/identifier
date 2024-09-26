package coding

import "strings"

type Alphabet string

func (a Alphabet) Char(offset, index int64) byte {
	return a[(offset+index)%int64(len(a))]
}

// func (a Alphabet) Index(offset int64, char byte) int64 {
// 	l := int64(len(a))
// 	return (int64(strings.IndexByte(string(a), char)) - offset + l) % l
// }

func (a Alphabet) Index(offset int64, char byte) int64 {
	l := int64(len(a))
	idx := int64(strings.IndexByte(string(a), char))

	// Adjusting to handle negative indices properly
	if idx == -1 {
		return -1 // Character not found in the alphabet (shouldn't happen)
	}

	adjustedIdx := (idx - offset) % l

	if adjustedIdx < 0 {
		adjustedIdx += l
	}

	return adjustedIdx
}
