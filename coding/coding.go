package coding

import "math"

func Encode(alpha Alphabet, v int64) string {
	buf := make([]byte, 0, encodingLength(alpha, v)+1)
	buf = encode(buf, alpha, 0, v)
	return b2s(buf)
}

func encode(buf []byte, alpha Alphabet, round, v int64) []byte {
	buf = buf[:0]

	if round > int64(len(alpha)) {
		return buf
	}

	off := round + v
	buf = append(buf, alpha.Char(round, v))
	buf = encodeVal(buf, alpha, off, v)

	return buf
}

func encodeVal(buf []byte, alpha Alphabet, offset, v int64) []byte {
	l := int64(len(alpha))

	for v != 0 {
		idx := v % l
		buf = append(buf, alpha.Char(offset, idx))
		v /= l
	}

	return buf
}

func Decode(alpha Alphabet, s string) int64 {
	return decode(s2b(s), alpha)
}

func decode(buf []byte, alpha Alphabet) (v int64) {
	off := alpha.Index(0, buf[0])
	return decodeVal(buf[1:], alpha, off)
}

func decodeVal(buf []byte, alpha Alphabet, offset int64) (v int64) {
	l := int64(len(alpha))
	for i := len(buf) - 1; i >= 0; i-- {
		char := buf[i]
		index := alpha.Index(offset, char)
		v = v*l + index
	}
	return v
}

func encodingLength(alpha Alphabet, v int64) int64 {
	l := int64(len(alpha))
	if v == 0 {
		return 1
	}
	return int64(math.Floor(math.Log(float64(v))/math.Log(float64(l)))) + 1
}
