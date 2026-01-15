package gotils

import (
	"maps"
	"numbers"
	"slices"
	"unicode/utf8"
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
			n, err := RandInt(int64(l - 1))

			if err != nil {
				return nil, err
			}

			i = int(n.Int64())
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

func encode(in string, rMap rMap, fMinNoise float32, fMaxNoise float32) (string, error) {
	out := make([]rune, len(in))
	noiseUnavailableRunes := make([]rune, 0)

	for i, r := range in {
		out[i] = rMap[r]
		noiseUnavailableRunes = append(noiseUnavailableRunes, rMap[r])
	}

	noiseAvailableRunes := Difference(slices.Sorted(maps.Values(rMap)), noiseUnavailableRunes)
	qNoiseAvailableRunes := float32(len(noiseAvailableRunes))
	minNoise := uint32(Clamp(qNoiseAvailableRunes*fMinNoise, 0, qNoiseAvailableRunes))
	maxNoise := uint32(Clamp(qNoiseAvailableRunes*fMaxNoise, 0, qNoiseAvailableRunes))

	if minNoise > 0 && maxNoise > 0 && maxNoise >= minNoise {
		steps := maxNoise - minNoise

		if steps > 0 {
			n, err := RandInt(int64(steps))

			if err != nil {
				return "", err
			}

			steps = uint32(n.Uint64())
		}

		for i := range minNoise + steps {
			j := 0
			
			if l := len(out); l > 0 {
				n, err := RandInt(int64(l))

				if err != nil {
					return "", err
				}

				j = int(n.Int64())
			}


			out = slices.Insert(out, j, noiseAvailableRunes[i])
		}

	}

	// slices.sor
	// len(reservedRunes)
	// slices
	// make([]rune, )

	// reservedRunes

}
