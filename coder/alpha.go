package coder

import (
	"math/rand"
	"strings"
)

const defaultAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ShuffleAlpha(seed int64, alpha ...string) string {
	var buf []byte

	if len(alpha) > 0 {
		buf = []byte(alpha[0])
	} else {
		buf = []byte(defaultAlpha)
	}

	rand.New(rand.NewSource(seed)).Shuffle(len(buf), func(i, j int) {
		buf[i], buf[j] = buf[j], buf[i]
	})

	return b2s(buf)
}

func (c *Coder) validateAlpha() (err error) {
	if len(c.alpha) < 3 {
		return ErrAlphaTooShort
	}

	if !uniqueAlpha(c.alpha) {
		return ErrAlphaNotUnique
	}

	return
}

func (c *Coder) alphaChar(offset, index int64) byte {
	return c.alpha[(offset+index)%c.alphaLen]
}

func (c *Coder) alphaIndex(offset int64, char byte) int64 {
	idx := int64(strings.IndexByte(c.alpha, char))

	// Adjusting to handle negative indices properly
	if idx == -1 {
		return -1 // Character not found in the alphabet (shouldn't happen)
	}

	adjustedIdx := (idx - offset) % c.alphaLen

	if adjustedIdx < 0 {
		adjustedIdx += c.alphaLen
	}

	return adjustedIdx
}

func uniqueAlpha(alpha string) bool {
	for i := 1; i < len(alpha); i++ {
		if strings.IndexByte(alpha[i:], alpha[i-1]) != -1 {
			return false
		}
	}

	return true
}
