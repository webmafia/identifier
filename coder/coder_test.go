package coder

import "fmt"

func Example() {
	c, err := NewCoder()
	// c, err := NewCoder("bcdfghjklmnpqrstvwxzBCDFGHJKLMNPQRSTVWXZ0123456789")

	if err != nil {
		panic(err)
	}

	v := int64(12345678)
	v2 := int64(12345679)
	enc := c.Encode(v)
	enc2 := c.Encode(v2)

	fmt.Println(v, "=", enc)
	fmt.Println(v2, "=", enc2)
	fmt.Println(c.Decode(enc), "(decoded)")

	// Output: TODO
}
