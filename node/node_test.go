package node

import (
	"fmt"
	"testing"
)

func BenchmarkIdentifier_bits(b *testing.B) {
	for i := uint8(22); i > 0; i-- {
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			n, err := NewNode(1, Options{
				NodeBits: i,
			})

			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()

			for range b.N {
				_ = n.Generate()
			}
		})
	}
}

func BenchmarkIdentifier_Parallell_bits(b *testing.B) {
	for i := uint8(22); i > 0; i-- {
		b.Run(fmt.Sprintf("%02d", i), func(b *testing.B) {
			n, err := NewNode(1, Options{
				NodeBits: i,
			})

			if err != nil {
				b.Fatal(err)
			}

			b.ResetTimer()

			b.RunParallel(func(p *testing.PB) {
				for p.Next() {
					_ = n.Generate()
				}
			})
		})
	}
}
