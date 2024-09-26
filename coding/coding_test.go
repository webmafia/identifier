package coding

import (
	"fmt"
	"testing"
)

func Example_encode() {
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var buf []byte

	a := int64(1839345119141068800)

	fmt.Println("length:", encodingLength(alpha, a))

	buf = encodeVal(buf, alpha, 72, a)
	fmt.Println("encoded:", string(buf))

	b := decodeVal(buf, alpha, 72)
	fmt.Println("decoded:", b)

	// Output:
	//
	// length: 11
	// encoded: GDa8xCYxcvm
	// decoded: 1839345119141068800
}

func ExampleEncode() {
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	a := int64(1026)
	fmt.Println(a, "=", Encode(alpha, a))

	// Output: TODO
}

func BenchmarkEncode(b *testing.B) {
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b.ResetTimer()

	for range b.N {
		_ = Encode(alpha, 1839345119141068800)
	}
}

func BenchmarkDecode(b *testing.B) {
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := Encode(alpha, 12345678)
	b.ResetTimer()

	for range b.N {
		_ = Decode(alpha, s)
	}
}
