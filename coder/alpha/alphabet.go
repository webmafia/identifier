package alpha

import (
	"math/rand"
	"strings"
)

type Alphabet struct {
	alphabet string
	offset   int64
	start    int64
	length   int64
}

// newAlphabet creates a new Alphabet without any initial rotation.
func NewAlphabet(charset ...string) Alphabet {
	var alpha string

	if len(charset) > 0 {
		alpha = charset[0]
	} else {
		alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	return Alphabet{
		alphabet: alpha,
		offset:   0,
		start:    0,
		length:   int64(len(alpha)),
	}
}

// Len returns the length of the sliced rotated alphabet.
func (r Alphabet) Len() int64 {
	return r.length
}

// At returns the byte at the given index in the sliced rotated alphabet.
func (r Alphabet) At(i int64) byte {
	if i < 0 || i >= r.length {
		panic("index out of range")
	}
	n := int64(len(r.alphabet))
	return r.alphabet[(r.offset+r.start+i)%n]
}

// Slice simulates slicing on the rotated alphabet.
func (r Alphabet) Slice(low, high int64) Alphabet {
	if low < 0 || high > r.length || low > high {
		panic("invalid slice indices")
	}
	return Alphabet{
		alphabet: r.alphabet,
		offset:   r.offset,
		start:    (r.start + low) % int64(len(r.alphabet)),
		length:   high - low,
	}
}

// Rotate rotates the alphabet by the given offset.
func (r Alphabet) Rotate(offset int64) Alphabet {
	n := int64(len(r.alphabet))
	offset = ((offset % n) + n) % n // Handle negative offsets
	return Alphabet{
		alphabet: r.alphabet,
		offset:   (r.offset + offset) % n,
		start:    r.start,
		length:   r.length,
	}
}

// IndexByte returns the index of the given byte in the rotated and sliced alphabet, or -1 if not found.
func (r Alphabet) IndexByte(b byte) int64 {
	n := int64(len(r.alphabet))
	originalIndex := int64(strings.IndexByte(r.alphabet, b))
	if originalIndex == -1 {
		return -1 // Character not found in the original alphabet
	}
	// Calculate the index within the rotated and sliced alphabet
	i := (originalIndex - r.offset - r.start + n) % n
	if i < 0 || i >= r.length {
		return -1 // Character not within the current view
	}
	return i
}

func (r Alphabet) Shuffle(seed int64) Alphabet {
	buf := []byte(r.alphabet)
	rand.New(rand.NewSource(seed)).Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return Alphabet{
		alphabet: string(buf),
		offset:   r.offset,
		start:    r.start,
		length:   r.length,
	}
}

func (r Alphabet) Validate() (err error) {
	if r.Len() < alphaMinChars {
		return ErrTooShort
	}

	for i, c := range r.alphabet {
		if strings.IndexByte(r.alphabet[i+1:], byte(c)) >= 0 {
			return ErrDuplicateChars
		}
	}

	return
}
