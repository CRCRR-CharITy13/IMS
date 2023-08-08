package utils

import (
	"crypto/rand"
	"math"
)

func GenerateProductId() int64 {
	// generate a random 8 digit number cryptographically securely

	p, _ := rand.Prime(rand.Reader, 64)

	return int64(math.Abs(float64(p.Int64())))
}
