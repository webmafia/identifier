package coder

import (
	"math"
)

func (c *Coder) Encode(v int64) string {
	buf := make([]byte, 0, c.minLen)
	buf = c.encode(buf, v)
	return b2s(buf)
}

func (c *Coder) EncodeToSlice(buf []byte, v int64) []byte {
	return c.encode(buf, v)
}

func (c *Coder) encode(buf []byte, v int64) []byte {
	off := v
	buf = append(buf, c.alphaChar(0, v))
	buf = c.encodeVal(buf, off, v)
	buf = c.padVal(buf, off)

	return buf
}

func (c *Coder) encodeVal(buf []byte, offset, v int64) []byte {
	for v != 0 {
		idx := v % c.alphaLen
		buf = append(buf, c.alphaChar(offset, idx))
		v /= c.alphaLen
	}

	return buf
}

func (c *Coder) padVal(buf []byte, offset int64) []byte {
	bufLen := int64(len(buf))

	if bufLen >= c.minLen {
		return buf
	}

	buf = append(buf, c.alphaChar(offset, 0))

	for bufLen < c.minLen {
		buf = append(buf, c.alphaChar(offset, bufLen*c.minLen))
		bufLen++
	}

	return buf
}

func encodingLength(alphaLen, maxVal int64) int64 {
	if maxVal == 0 {
		return 1 // Special case: 0 requires at least one character
	}

	// Perform the logarithmic calculation
	length := math.Ceil(math.Log(float64(maxVal)) / math.Log(float64(alphaLen)))

	// Check if an additional character is needed by comparing actual encoding range
	if int64(math.Pow(float64(alphaLen), length)) <= maxVal {
		length++
	}

	return int64(length)
}
