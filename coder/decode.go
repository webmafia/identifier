package coder

func (c *Coder) Decode(s string) int64 {
	return c.decode(s2b(s))
}

func (c *Coder) decode(buf []byte) (v int64) {
	off := c.alphaIndex(0, buf[0])
	return c.decodeVal(buf[1:], off)
}

func (c *Coder) decodeVal(buf []byte, offset int64) (v int64) {
	multiplier := int64(1)
	for i := 0; i < len(buf); i++ {
		char := buf[i]
		index := c.alphaIndex(offset, char)

		if index == 0 {
			break
		}

		v += index * multiplier
		multiplier *= c.alphaLen
	}

	return v
}
