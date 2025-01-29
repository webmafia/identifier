package identifier

func FromString(str string) (id ID, err error) {
	v, err := cod.Decode(str)

	if err != nil {
		return
	}

	return ID(v), nil
}
