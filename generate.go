package identity

// Generates a unique ID. This is thread-safe.
func GenerateId() ID {
	return ID(gen.Generate())
}
