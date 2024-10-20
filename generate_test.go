package identifier

import (
	"fmt"
	"math"
	"testing"
)

func ExampleGenerate() {
	for range 10 {
		id := Generate()
		fmt.Println(id.Int64(), id)
	}

	// Output: TODO
}

func BenchmarkGenerate(b *testing.B) {
	b.Run("SingleThread", func(b *testing.B) {
		for range b.N {
			_ = Generate()
		}
	})

	b.Run("Parallell", func(b *testing.B) {
		b.RunParallel(func(p *testing.PB) {
			for p.Next() {
				_ = Generate()
			}
		})
	})
}

func BenchmarkStringer(b *testing.B) {
	id := ID(math.MaxInt64)
	b.ResetTimer()

	for range b.N {
		_ = id.String()
	}
}
