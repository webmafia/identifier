package format

import (
	"fmt"
	"strings"
	"testing"
)

func Example() {
	n, err := New()

	if err != nil {
		panic(err)
	}

	v1 := int64(12345678)
	v2 := int64(87654321)
	enc1, _ := n.Encode(v1)
	enc2, _ := n.Encode(v2)

	fmt.Println(enc1, n.Decode(enc1))
	fmt.Println(enc2, n.Decode(enc2))

	// Output:
	//
	// QZ9JU 12345678
	// roVNwl 87654321
}

func Benchmark(b *testing.B) {
	n, err := New()

	if err != nil {
		panic(err)
	}

	for range b.N {
		_, _ = n.Encode(12345679)
	}
}

func Example_toID() {
	var b strings.Builder
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	toID(12345678, alpha, &b)

	fmt.Println(b.String())

	// Output: TODO
}

func Example_toNumber() {
	var b strings.Builder
	alpha := Alphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	toID(12345678, alpha, &b)
	str := b.String()

	num := toNumber([]byte(str), alpha)

	fmt.Println(b.String())
	fmt.Println(num)

	// Output: TODO
}
