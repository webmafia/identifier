package identifier

// Generates a unique ID. This is thread-safe.
func Generate() ID {
	return ID(gen.Generate())
}
