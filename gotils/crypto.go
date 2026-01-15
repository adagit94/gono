package gotils

import (
	"numbers"
	"slices"
	"maps"
	"unicode/utf8"
	"math/big"
	"crypto/rand"
)

// const (
// 	n0, n9   uint8 = 48, 57
// 	aUp, zUp uint8 = 65, 90
// 	aLo, zLo uint8 = 97, 122
// )

type rMap map[rune]rune
type rOffset uint16

const runesCount = utf8.MaxRune + 1

func GenQMaps() (rMap, error) {
	s := make([]rune, runesCount)

	for i := range s {
		s[i] = rune(i)
	}

	m := make(rMap, runesCount)

	for r := range runesCount {
		i := 0

		if l := len(s); l > 1 {
			v, err := rand.Int(rand.Reader, big.NewInt(int64(l-1)))

			if err != nil {
				return nil, err
			}

			i = int(v.Int64())
		}

		m[r] = s[i]
		s = slices.Delete(s, i, i+1)
	}

	return m, nil
}

type key struct {
	rOffset rOffset
	rMap    rMap
}

func GenKey(rMap rMap, rOffset rOffset) key {

}

func encode(in string, rMap rMap, minNoiseF uint8, maxNoiseF uint8) string {
	reservedRunes := slices.Sorted(maps.Values(rMap))
	availableRunes := make([]rune, runesCount - len(reservedRunes))
	//  outS := make([]rune, )

	 for r := range in {

	 }

	slices.sor
	// len(reservedRunes)
	// slices
	// make([]rune, )

	// reservedRunes
}
