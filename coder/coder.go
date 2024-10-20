package coder

import (
	"bytes"
	"math"

	"github.com/webmafia/identifier/coder/alpha"
)

// Number encoder
type Coder struct {
	alpha  alpha.Alphabet
	encLen int
}

func New(a ...alpha.Alphabet) (*Coder, error) {
	var alp alpha.Alphabet

	if len(a) > 0 {
		alp = a[0]
	} else {
		alp = alpha.NewAlphabet().Shuffle(1337)
	}

	if err := alp.Validate(); err != nil {
		return nil, err
	}

	return &Coder{
		alpha:  alp,
		encLen: encodingLength(alp.Len(), math.MaxInt64),
	}, nil
}

func (c *Coder) Encode(n int64) string {
	buf := make([]byte, c.encLen)
	c.encode(buf, n)
	return b2s(buf)
}

func (c *Coder) AppendEncoded(buf []byte, n int64) []byte {
	if l := len(buf); cap(buf) >= l+c.encLen {
		buf = buf[:l+c.encLen]
	} else {
		newBuf := make([]byte, l+c.encLen)
		copy(newBuf, buf)
		buf = newBuf
	}

	c.encode(buf[len(buf)-c.encLen:], n)

	return buf
}

func (c *Coder) Decode(s string) (n int64, err error) {
	return c.DecodeBytes(s2b(s))
}

func (c *Coder) DecodeBytes(b []byte) (n int64, err error) {
	if len(b) != c.encLen {
		return 0, ErrInvalidString
	}

	return c.decode(b)
}

func (c *Coder) encode(dst []byte, n int64) {
	_ = dst[7]

	offset := c.calculateOffset(n)
	alpha := c.alpha.Rotate(offset)
	head := alpha.At(0)

	alpha = alpha.Slice(1, alpha.Len())
	result := n
	i := int64(len(dst) - 1)
	dst[i] = head
	i--

	for i >= 0 && result > 0 {
		index := result % alpha.Len()
		dst[i] = alpha.At(index)
		result = result / alpha.Len()

		i--
	}

	if i >= 0 {
		dst[i] = head
		i--

		for i >= 0 {
			index := (n*i*int64(head) + i*1993) % alpha.Len()
			dst[i] = alpha.At(index)

			i--
		}
	}
}

// Decode id string into a slice of int64 values
func (c *Coder) decode(src []byte) (n int64, err error) {
	_ = src[:c.encLen]

	head := src[c.encLen-1]
	offset := c.alpha.IndexByte(head)

	if offset < 0 {
		return 0, ErrInvalidString
	}

	src = src[:c.encLen-1]

	if i := bytes.IndexByte(src, head); i > 0 {
		src = src[i+1:]
	}

	alpha := c.alpha.Rotate(offset).Slice(1, c.alpha.Len())
	count := alpha.Len()

	for _, r := range src {
		idx := alpha.IndexByte(r)

		if idx < 0 {
			return 0, ErrInvalidString
		}

		n = (n * count) + idx
	}

	return
}

func (c *Coder) calculateOffset(n int64) int64 {
	return n * 1993 % c.alpha.Len()
}

func (c *Coder) EncodedLength() int {
	return c.encLen
}
