package coder

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func Example() {
	s, err := New()

	if err != nil {
		panic(err)
	}

	for i := range int64(10) {
		enc := s.Encode(i)
		dec, _ := s.Decode(enc)

		fmt.Printf("%02d = %s = %02d\n", i, enc, dec)
	}

	fmt.Println(strings.Repeat("-", 24))

	for i := int64(math.MaxInt64 - 10); i < math.MaxInt64; i++ {
		enc := s.Encode(i)
		dec, _ := s.Decode(enc)

		fmt.Printf("%02d = %s = %02d\n", i, enc, dec)
	}

	// Output:
	// 00 = 2ontTmrhpq11 = 00
	// 01 = JIEAcqrt2PfP = 01
	// 02 = VNHaJIAqtFQF = 02
	// 03 = KOEzks0FIili = 03
	// 04 = RTkGMcfOmbXb = 04
	// 05 = 6cejEKvP39x9 = 05
	// 06 = GZzM2rcEPUMU = 06
	// 07 = ti5NXc3GV2P2 = 07
	// 08 = fwZVorivaJVJ = 08
	// 09 = CotWDE35aV0V = 09
	// ------------------------
	// 9223372036854775797 = 6KSOdrosiS37 = 9223372036854775797
	// 9223372036854775798 = GRpehaxjbpyo = 9223372036854775798
	// 9223372036854775799 = t67dYQuZ97vx = 9223372036854775799
	// 9223372036854775800 = fGoh13EqUTMu = 9223372036854775800
	// 9223372036854775801 = CtxYPsN825PE = 9223372036854775801
	// 9223372036854775802 = 0fu1FjWnJMVN = 9223372036854775802
	// 9223372036854775803 = BCEPiZwlVI0W = 9223372036854775803
	// 9223372036854775804 = D0NFbqTyKOsw = 9223372036854775804
	// 9223372036854775805 = LBWi985kRekT = 9223372036854775805
	// 9223372036854775806 = rDwbUnM46dz5 = 9223372036854775806
}

func BenchmarkEncode(b *testing.B) {
	s, err := New()

	if err != nil {
		panic(err)
	}

	b.ResetTimer()

	for i := range int64(b.N) {
		_ = s.Encode(i)
	}
}

func BenchmarkAppendEncoded(b *testing.B) {
	s, err := New()

	if err != nil {
		panic(err)
	}

	var buf []byte

	b.ResetTimer()

	for i := range int64(b.N) {
		buf = s.AppendEncoded(buf[:0], i)
	}
}

func BenchmarkDecode(b *testing.B) {
	s, err := New()

	if err != nil {
		panic(err)
	}

	v := s.Encode(math.MaxInt64)

	b.ResetTimer()

	for range b.N {
		_, _ = s.Decode(v)
	}
}

func Benchmark_encodingLength(b *testing.B) {
	for i := range int64(b.N) {
		_ = encodingLength(62, i)
	}
}
