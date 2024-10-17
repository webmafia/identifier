package coder

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

func Test_encodingLength(t *testing.T) {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := len(alpha); i >= 3; i-- {
		t.Run(fmt.Sprintf("alpha%02d", i), func(t *testing.T) {
			c, err := NewCoder(alpha[:i])

			if err != nil {
				t.Fatal(err)
			}

			if i != int(c.alphaLen) {
				t.Errorf("expected alphaLen to be %d, but got %d", i, c.alphaLen)
			}

			t.Logf("alphaLen: %d, maxVal: %d", c.alphaLen, math.MaxInt64)

			logMaxVal := math.Log(float64(math.MaxInt64))
			logAlphaLen := math.Log(float64(c.alphaLen))

			t.Logf("Log of maxVal: %f", logMaxVal)
			t.Logf("Log of alphaLen: %f", logAlphaLen)

			// Calculate predicted length
			predictedLen := int(encodingLength(c.alphaLen, math.MaxInt64))
			t.Logf("Predicted length: %d", predictedLen)

			var buf []byte
			buf = c.encode(buf, math.MaxInt64)

			if l := len(buf); l != predictedLen {
				t.Errorf("predicted %d, but turned out to be %d", predictedLen, l)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	alpha := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	c, err := NewCoder(alpha)

	if err != nil {
		t.Fatal(err)
	}

	// Predicted length based on the maximum possible value
	predictedLen := int(c.EncodedLength())
	rand := rand.New(rand.NewSource(1337))

	for i := 0; i < 10_000; i++ {
		v := rand.Int63()
		s := c.Encode(v)

		// Adjust if the encoded length exceeds the prediction (account for edge cases)
		if l := len(s); l > predictedLen {
			t.Errorf("%d: predicted %d bytes, but got %d bytes. Value: %d", i, predictedLen, l, v)
			t.FailNow()
		}
	}
}
