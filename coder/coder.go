package coder

import (
	"errors"
	"math"
)

var (
	ErrAlphaTooShort  = errors.New("alphabet is too short (must be at least 3 characters)")
	ErrAlphaNotUnique = errors.New("not unique characters (bytes) in alphabet")
	ErrNotInAlpha     = errors.New("encountered character that is not in alphabet")
)

type Coder struct {
	alpha    string
	alphaLen int64
	minLen   int64
}

func NewCoder(alpha ...string) (c *Coder, err error) {
	c = new(Coder)

	if len(alpha) > 0 {
		c.alpha = alpha[0]
	} else {
		c.alpha = ShuffleAlpha(1337)
	}

	if err = c.validateAlpha(); err != nil {
		return nil, err
	}

	c.alphaLen = int64(len(c.alpha))
	c.minLen = c.encodingLength(math.MaxInt64)

	return
}

func (c *Coder) EncodedLength() int64 {
	return c.minLen
}
