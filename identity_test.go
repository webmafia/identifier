package identity

import "testing"

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
