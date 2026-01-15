package gotils

import (
	"cmp"
	"crypto/rand"
	"math/big"
)

func Clamp[T cmp.Ordered](v T, min T, max T) T {
	if v < min {
		return min
	}

	if v > max {
		return max
	}

	return v
}

func RandInt(max int64) (n *big.Int, err error) {
	return rand.Int(rand.Reader, big.NewInt(max))
}
