package numbers

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

func RandInt(max int64) (*big.Int, error) {
	return rand.Int(rand.Reader, big.NewInt(max))
}

func RandIntPanic(max int64) *big.Int {
	n, err := RandInt(max)

	if err != nil {
		panic(err)
	}

	return n
}
