package identity

type ID int64

// MarshalJSON returns a json byte array string of the snowflake ID.
func (id ID) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

// UnmarshalJSON converts a json byte array of a snowflake ID into an ID type.
func (id *ID) UnmarshalJSON(b []byte) error {
	panic("not implemented")
}
