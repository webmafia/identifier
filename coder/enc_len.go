package coder

import "math"

func encodingLength(alphaLen, maxVal int64) int {
	if maxVal == 0 {
		return 1 // Special case: 0 requires at least one character
	}

	// Perform the logarithmic calculation
	length := math.Ceil(math.Log(float64(maxVal)) / math.Log(float64(alphaLen)))

	// Check if an additional character is needed by comparing actual encoding range
	if int64(math.Pow(float64(alphaLen), length)) <= maxVal {
		length++
	}

	return int(length)
}
