package coder

import (
	"fmt"
	"math"
)

func Example() {
	c, err := NewCoder()

	if err != nil {
		panic(err)
	}

	v := int64(12345678)
	v2 := int64(12345679)
	enc := c.Encode(v)
	enc2 := c.Encode(v2)

	fmt.Println(v, "=", enc)
	fmt.Println(v2, "=", enc2)
	fmt.Println(c.Decode(enc))

	// Output:
	//
	// 12345678 = 5olsX5v1apbju
	// 12345679 = YhgyoYx28WRkM
	// 12345678 <nil>
}

func Example_minMax() {
	c, err := NewCoder()

	if err != nil {
		panic(err)
	}

	min := int64(0)
	minEnc := c.Encode(min)

	max := int64(math.MaxInt64)
	maxEnc := c.Encode(max)

	fmt.Printf("%d = %s\n", min, minEnc)
	fmt.Printf("%d = %s\n", max, maxEnc)
	fmt.Println(c.Decode(minEnc))
	fmt.Println(c.Decode(maxEnc))

	// Output:
	//
	// 9223372036854775807 = EusM4WEyMa52E
	// 9223372036854775807 <nil>
}
