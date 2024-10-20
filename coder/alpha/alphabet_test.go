package alpha

import (
	"fmt"
	"testing"
)

func ExampleAlphabet() {
	alpha := NewAlphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789").Rotate(500)
	// alpha = alpha.Slice(1, 10)

	for i := range alpha.Len() {
		fmt.Printf("%02d = %s\n", i, string(alpha.At(i)))
	}

	fmt.Println("found 'd' at:", alpha.IndexByte('d'))

	// Output: TODO
}

func BenchmarkAlphabetRotate(b *testing.B) {
	alpha := NewAlphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b.ResetTimer()

	for range b.N {
		alpha = alpha.Rotate(1)
	}
}

func BenchmarkAlphabetIndexByte(b *testing.B) {
	alpha := NewAlphabet("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b.ResetTimer()

	for range b.N {
		_ = alpha.IndexByte('D')
	}
}
