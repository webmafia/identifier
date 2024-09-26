package format

import (
	"errors"
	"strings"
)

const (
	defaultAlphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	minAlphabetLength = 3
)

// Alphabet represents the alphabet used for encoding/decoding
type Alphabet string

// Reverse returns a new Alphabet instance with the characters reversed
func (a Alphabet) Reverse() Alphabet {
	size := len(a)
	reversed := make([]byte, size)

	// Reverse the characters by iterating from both ends
	for i := 0; i < size; i++ {
		reversed[i] = a[size-1-i]
	}

	// Return a new Alphabet instance from the reversed string
	return Alphabet(reversed)
}

// OffsetAndAppend appends the alphabet from the given offset and length directly to the Builder
func (a Alphabet) OffsetAndAppend(offset int, length int, b *strings.Builder) {
	size := len(a)

	// Handle cases with wrap-around
	if offset+length <= size {
		// No wrap-around, append directly
		b.WriteString(string(a[offset : offset+length]))
	} else {
		// Wrap-around: append from offset to end, then from start
		b.WriteString(string(a[offset:]))
		b.WriteString(string(a[:length-(size-offset)]))
	}
}

// Alphabet validation errors
var (
	errAlphabetMultibyte       = errors.New("alphabet must not contain any multibyte characters")
	errAlphabetTooShort        = errors.New("alphabet length must be at least 3")
	errAlphabetNotUniqueChars  = errors.New("alphabet must contain unique characters")
	errMaxRegenerationAttempts = errors.New("reached max attempts to re-generate the id")
)

// Options for a custom instance of Sqids
type Options struct {
	Alphabet  string
	MinLength uint8
}

// Sqids lets you generate unique IDs from a single number
type Sqids struct {
	alphabet        Alphabet
	reverseAlphabet Alphabet
	minLength       uint8
}

// New constructs an instance of Sqids
func New(options ...Options) (*Sqids, error) {
	if len(options) == 0 {
		options = append(options, Options{
			Alphabet: defaultAlphabet,
		})
	}

	// Validate the first given options value, or the default options if none were given.
	o, err := validatedOptions(options[0])
	if err != nil {
		return nil, err
	}

	// Initialize the alphabet and reverseAlphabet once
	alphabet := Alphabet(o.Alphabet)
	reverseAlphabet := alphabet.Reverse()

	return &Sqids{
		alphabet:        alphabet,
		reverseAlphabet: reverseAlphabet,
		minLength:       o.MinLength,
	}, nil
}

func validatedOptions(o Options) (Options, error) {
	if o.Alphabet == "" {
		o.Alphabet = defaultAlphabet
	}

	// check that the alphabet does not contain multibyte characters
	if len(o.Alphabet) != len([]byte(o.Alphabet)) {
		return Options{}, errAlphabetMultibyte
	}

	// check the length of the alphabet
	if len(o.Alphabet) < minAlphabetLength {
		return Options{}, errAlphabetTooShort
	}

	// check that the alphabet has only unique characters
	if !hasUniqueChars([]byte(o.Alphabet)) {
		return Options{}, errAlphabetNotUniqueChars
	}

	return o, nil
}

// Encode a single int64 value into an ID string
func (s *Sqids) Encode(number int64) (string, error) {
	// If the number is zero or less, return an empty string
	if number <= 0 {
		return "", nil
	}

	return s.encodeNumber(number, 0)
}

func (s *Sqids) encodeNumber(number int64, increment int) (string, error) {
	if increment > len(s.alphabet) {
		return "", errMaxRegenerationAttempts
	}

	var (
		offset = calculateOffset(s.alphabet, number, increment)
		b      strings.Builder
	)

	b.Grow(int(s.minLength))
	s.alphabet.OffsetAndAppend(offset, 1, &b) // Write the prefix

	// Append the encoded number using the reversed alphabet
	toID(number, s.reverseAlphabet, &b)

	// If the ID is shorter than the minimum length, pad it
	if int(s.minLength) > b.Len() {
		for int(s.minLength)-b.Len() > 0 {
			s.alphabet.OffsetAndAppend(0, min(int(s.minLength)-b.Len(), len(s.alphabet)), &b)
		}
	}

	return b.String(), nil
}

// Decode an ID string into a single int64 value
func (s *Sqids) Decode(id string) int64 {
	if id == "" {
		return 0
	}

	rid := []byte(id)
	alphabet := []byte(s.alphabet)

	// The first character (prefix) determines the offset
	prefix := rid[0]
	offset := index(alphabet, prefix)

	// Adjust the reverseAlphabet based on the offset
	adjustedAlphabet := s.reverseAlphabet
	adjustedAlphabet.OffsetAndAppend(offset, len(adjustedAlphabet), &strings.Builder{}) // Adjust the reverse alphabet using the offset

	// Decode the rest of the string using the adjusted reverse alphabet
	return toNumber(rid[1:], adjustedAlphabet)
}

// toID encodes the number into the alphabet
func toID(num int64, alphabet Alphabet, b *strings.Builder) {
	count := int64(len(alphabet))
	result := num

	for {
		index := result % count
		b.WriteByte(alphabet[index])
		result = result / count

		if result == 0 {
			break
		}
	}
}

// toNumber decodes the string into the original number
func toNumber(id []byte, alphabet Alphabet) int64 {
	count := int64(len(alphabet))
	var result int64

	// Decode the ID from the alphabet
	for _, ch := range id {
		result = (result * count) + int64(index([]byte(alphabet), ch))
	}

	return result
}

func calculateOffset(alphabet Alphabet, number int64, increment int) int {
	count := len(alphabet)
	offset := int(number%int64(count)) + increment
	return offset % count
}

func index(s []byte, b byte) int {
	for i := range s {
		if b == s[i] {
			return i
		}
	}
	return -1
}

func hasUniqueChars(s []byte) bool {
	charSet := make(map[byte]bool)
	for _, b := range s {
		if _, ok := charSet[b]; ok {
			return false
		}
		charSet[b] = true
	}
	return true
}

func contains(s []byte, b byte) bool {
	for _, v := range s {
		if v == b {
			return true
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
