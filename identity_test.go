package identity

import (
	"fmt"
	"testing"

	"github.com/sqids/sqids-go"
)

func Example() {
	coder, err := sqids.New(sqids.Options{
		Alphabet:  "GoRxYqtn02dfHCT8SFeVBgrzviO49WLUmuQj5pywcKPZ3lb7kMAJaE6NDIhXs1",
		MinLength: 12,
	})

	if err != nil {
		panic(err)
	}

	n, err := NewNode(1, Options{
		Coder: coder,
	})

	if err != nil {
		panic(err)
	}

	id := n.Generate()
	fmt.Println(id, "=", n.ToString(id))

	// Output: TODO
}

func BenchmarkIdentityCoder(b *testing.B) {
	coder, err := sqids.New(sqids.Options{
		Alphabet:  "GoRxYqtn02dfHCT8SFeVBgrzviO49WLUmuQj5pywcKPZ3lb7kMAJaE6NDIhXs1",
		MinLength: 12,
	})

	if err != nil {
		panic(err)
	}

	n, err := NewNode(1, Options{
		Coder: coder,
	})

	if err != nil {
		panic(err)
	}

	id := n.Generate()
	b.ResetTimer()

	for range b.N {
		_ = n.ToString(id)
	}
}

func BenchmarkIdentity(b *testing.B) {
	n, err := NewNode(1)

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for range b.N {
		_ = n.Generate()
	}
}

func BenchmarkIdentity_Parallell(b *testing.B) {
	n, err := NewNode(1)

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = n.Generate()
		}
	})
}

func BenchmarkIdentity_bits(b *testing.B) {
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

func BenchmarkIdentity_Parallell_bits(b *testing.B) {
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
