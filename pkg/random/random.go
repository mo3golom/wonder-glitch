package random

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	offset := 0
	input := max - min

	// Intn hates 0 or less, so we use this workaround
	if input <= 0 {
		offset = 1 + input*-1
		input = offset
	}

	return RandomWithSeed(min, max, time.Now().UnixNano())
}

func RandomWithSeed(min, max int, seed int64) int {
	offset := 0
	input := max - min

	// Intn hates 0 or less, so we use this workaround
	if input <= 0 {
		offset = 1 + input*-1
		input = offset
	}

	randomSource := rand.NewSource(seed)

	return rand.New(randomSource).Intn(input) + min - offset
}
