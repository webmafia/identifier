package identifier

import (
	"database/sql"
	"database/sql/driver"
	"encoding"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"time"
)

var (
	_ fmt.Stringer               = (*ID)(nil)
	_ json.Marshaler             = (*ID)(nil)
	_ json.Unmarshaler           = (*ID)(nil)
	_ encoding.BinaryMarshaler   = (*ID)(nil)
	_ encoding.BinaryUnmarshaler = (*ID)(nil)
	_ encoding.TextMarshaler     = (*ID)(nil)
	_ encoding.TextUnmarshaler   = (*ID)(nil)
	_ sql.Scanner                = (*ID)(nil)
	_ driver.Valuer              = (*ID)(nil)
	_ TextAppender               = (*ID)(nil)
	_ BinaryAppender             = (*ID)(nil)
)

type ID int64

// String implements fmt.Stringer.
func (id ID) String() string {
	return cod.Encode(int64(id))
}

func (id ID) Time() time.Time {
	return gen.Time(int64(id))
}

func (id ID) Int64() int64 {
	return int64(id)
}

func (id ID) IsZero() bool {
	return id == 0
}

func (id ID) IsNil() bool {
	return id == 0
}

// MarshalJSON implements json.Marshaler.
func (id ID) MarshalJSON() ([]byte, error) {
	buf := make([]byte, 0, cod.EncodedLength()+2)
	buf = append(buf, '"')
	buf = cod.AppendEncoded(buf, int64(id))
	buf = append(buf, '"')

	return buf, nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (id *ID) UnmarshalJSON(b []byte) (err error) {
	if len(b) == 4 && string(b) == "null" {
		*id = 0
		return
	}

	l := len(b)

	if l < 2 || b[0] != '"' || b[l-1] != '"' {
		return ErrInvalidId
	}

	b = b[1 : l-1]

	if len(b) == 0 {
		*id = 0
		return
	}

	v, err := cod.DecodeBytes(b)

	if err != nil {
		return
	}

	*id = ID(v)
	return
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (id ID) MarshalBinary() (data []byte, err error) {
	return id.AppendBinary(make([]byte, 0, 8))
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (id *ID) UnmarshalBinary(data []byte) error {
	if len(data) != 8 {
		return ErrInvalidId
	}

	*id = ID(binary.BigEndian.Uint64(data))
	return nil
}

// MarshalText implements encoding.TextMarshaler.
func (id ID) MarshalText() (text []byte, err error) {
	return id.AppendText(make([]byte, 0, cod.EncodedLength()))
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (id *ID) UnmarshalText(text []byte) (err error) {
	v, err := cod.DecodeBytes(text)

	if err != nil {
		return
	}

	*id = ID(v)
	return
}

// Scan implements sql.Scanner.
func (id *ID) Scan(src any) (err error) {
	switch v := src.(type) {
	case int64:
		*id = ID(v)
	case nil:
		*id = 0
	default:
		err = fmt.Errorf("cannot scan %T to %T", v, id)
	}

	return
}

// Value implements driver.Valuer.
func (id ID) Value() (driver.Value, error) {
	return int64(id), nil
}

// AppendText implements encoding.TextAppender.
func (id ID) AppendText(b []byte) ([]byte, error) {
	return cod.AppendEncoded(b, int64(id)), nil
}

// AppendBinary implements encoding.BinaryAppender.
func (id ID) AppendBinary(b []byte) ([]byte, error) {
	return binary.BigEndian.AppendUint64(b, uint64(id)), nil
}
