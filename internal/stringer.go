package internal

import (
	"errors"
	"strings"
)

const (
	defaultAlphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	minAlphabetLength = 3
)

var defaultBlocklist []string = newDefaultBlocklist()

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
	Blocklist []string
}

// Sqids lets you generate unique IDs from numbers
type Sqids struct {
	alphabet  string
	minLength uint8
	blocklist []string
}

// New constructs an instance of Sqids
func New(options ...Options) (*Sqids, error) {
	if len(options) == 0 {
		options = append(options, Options{
			Alphabet:  defaultAlphabet,
			Blocklist: defaultBlocklist,
		})
	}

	// Validate the first given options value, or the default options if none were given.
	o, err := validatedOptions(options[0])
	if err != nil {
		return nil, err
	}

	return &Sqids{
		alphabet:  shuffle(o.Alphabet),
		minLength: o.MinLength,
		blocklist: o.Blocklist,
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

	o.Blocklist = filterBlocklist(o.Alphabet, o.Blocklist)

	return o, nil
}

// Encode a slice of uint64 values into an ID string
func (s *Sqids) Encode(v int64) (string, error) {
	return s.encodeNumbers(v, 0)
}

func (s *Sqids) encodeNumbers(v int64, increment int) (string, error) {
	if increment > len(s.alphabet) {
		return "", errMaxRegenerationAttempts
	}

	var (
		err      error
		offset   = calculateOffset(s.alphabet, v, increment)
		alphabet = alphabetOffset(s.alphabet, offset)
		prefix   = alphabet[0]
		id       strings.Builder
	)

	id.Grow(int(s.minLength))

	alphabet = reversebytes(alphabet)

	ret = append(ret, []byte(toID(num, string(alphabet[1:])))...)

	id := string(ret)

	if int(s.minLength) > len(id) {
		id += string(alphabet[0])

		for int(s.minLength)-len(id) > 0 {
			alphabet = []byte(shuffle(string(alphabet)))
			id += string(alphabet[:min(int(s.minLength)-len(id), len(alphabet))])
		}
	}

	if s.isBlockedID(id) {
		id, err = s.encodeNumbers(numbers, increment+1)
		if err != nil {
			return "", err
		}
	}

	return id, nil
}

// Decode id string into a slice of uint64 values
func (s *Sqids) Decode(id string) []uint64 {
	ret := []uint64{}

	if id == "" {
		return ret
	}

	rid := []byte(id)

	alphabet := []byte(s.alphabet)

	for _, r := range rid {
		if !contains(alphabet, r) {
			return ret
		}
	}

	prefix := rid[0]
	offset := index(alphabet, prefix)

	alphabet = alphabetOffset(s.alphabet, offset)
	alphabet = reversebytes(alphabet)

	rid = rid[1:]

	for len(rid) > 0 {
		separator := alphabet[0]

		chunks := splitChunks(rid, separator)
		if len(chunks) > 0 {
			if len(chunks[0]) == 0 {
				return ret
			}

			ret = append(ret, toNumber(chunks[0], alphabet[1:]))

			if len(chunks) > 1 {
				alphabet = shufflebytes(alphabet)
			}
		}

		if len(chunks) > 0 {
			rid = joinbyteSlices(chunks[1:], separator)
		} else {
			return []uint64{}
		}
	}

	return ret
}

func alphabetOffset(alphabet string, offset int) []byte {
	bytes := []byte(alphabet)

	return append(bytes[offset:], bytes[:offset]...)
}

func joinbyteSlices(rs [][]byte, separator byte) []byte {
	var bytes []byte

	if len(rs) > 0 {
		for _, s := range rs[:len(rs)-1] {
			bytes = append(bytes, s...)
			bytes = append(bytes, separator)
		}

		bytes = append(bytes, rs[len(rs)-1]...)
	}

	return bytes
}

func splitChunks(bytes []byte, separator byte) [][]byte {
	var chunks [][]byte
	chunk := []byte{}

	for _, r := range bytes {
		if r == separator {
			chunks = append(chunks, chunk)
			chunk = []byte{}
		} else {
			chunk = append(chunk, r)
		}
	}

	chunks = append(chunks, chunk)
	return chunks
}

func (s *Sqids) isBlockedID(id string) bool {
	id = strings.ToLower(id)

	for _, word := range s.blocklist {
		if len(word) <= len(id) {
			if len(id) <= 3 || len(word) <= 3 {
				if id == word {
					return true
				}
			} else if hasDigit(word) {
				if strings.HasPrefix(id, word) || strings.HasSuffix(id, word) {
					return true
				}
			} else if strings.Contains(id, word) {
				return true
			}
		}
	}

	return false
}

func calculateOffset(alphabet string, v int64, increment int) int {
	offset := int(alphabet[v])
	offset = offset % len(alphabet)
	return (offset + increment) % len(alphabet)
}

func shuffle(alphabet string) string {
	return string(shufflebytes([]byte(alphabet)))
}

func shufflebytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; j > 0; i, j = i+1, j-1 {
		r := (i*j + int(bytes[i]) + int(bytes[j])) % len(bytes)
		bytes[i], bytes[r] = bytes[r], bytes[i]
	}

	return bytes
}

func reversebytes(bytes []byte) []byte {
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	return bytes
}

func toID(num uint64, alphabet string) string {
	var (
		id     = []byte{}
		bytes  = []byte(alphabet)
		count  = uint64(len(bytes))
		result = num
	)

	for {
		index := result % count

		id = append([]byte{bytes[index]}, id...)

		result = result / count

		if result == 0 {
			break
		}
	}

	return string(id)
}

func toNumber(rid []byte, bytes []byte) uint64 {
	count := uint64(len(bytes))

	var result uint64

	for _, r := range rid {
		result = (result * count) + uint64(index(bytes, r))
	}

	return result
}

func index(s []byte, r byte) int {
	for i := range s {
		if r == s[i] {
			return i
		}
	}

	return -1
}

func contains(s []byte, r byte) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}

	return false
}

func hasDigit(word string) bool {
	for _, r := range word {
		if r >= '0' && r <= '9' {
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
