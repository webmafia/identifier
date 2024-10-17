package coder

import (
	"math"
)

func (c *Coder) Encode(v int64) string {
	buf := make([]byte, 0, c.minLen)
	buf = c.encode(buf, 0, v)
	return b2s(buf)
}

func (c *Coder) EncodeToSlice(buf []byte, v int64) []byte {
	return c.encode(buf, 0, v)
}

func (c *Coder) encode(buf []byte, round, v int64) []byte {
	if round > c.alphaLen {
		return buf
	}

	off := round + v
	buf = append(buf, c.alphaChar(round, v))
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

func (c *Coder) encodingLength(v int64) int64 {
	if v == 0 {
		return 1
	}

	return int64(math.Floor(math.Log(float64(v))/math.Log(float64(c.alphaLen)))) + 2
}
