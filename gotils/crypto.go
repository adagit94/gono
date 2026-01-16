package gotils

import (
	"maps"
	"math"
	"numbers"
	"slices"
	"unicode/utf8"
)

// const (
// 	n0, n9   uint8 = 48, 57
// 	aUp, zUp uint8 = 65, 90
// 	aLo, zLo uint8 = 97, 122
// )

const runesCount = utf8.MaxRune + 1

type runesMap map[rune]rune

func GenRunesMaps(runesRangeFrom rune, runesRangeTo rune) (runesMap, runesMap, error) {
	runesCount := runesRangeTo - runesRangeFrom + 1
	s := make([]rune, runesCount)

	for i, v := 0, runesRangeFrom; v <= runesRangeTo; v++ {
		s[i] = v
		i++
	}

	m := make(runesMap, runesCount)
	m2 := make(runesMap, runesCount)

	for r := range runesCount {
		i := 0

		if l := len(s); l > 1 {
			n, err := RandInt(int64(l))

			if err != nil {
				return nil, nil, err
			}

			i = int(n.Uint64())
		}

		m[r] = s[i]
		m2[s[i]] = r  
		s = slices.Delete(s, i, i+1)
	}

	return m, m2, nil
}

func computeRunesRanges(fNoise float64) (rune, rune, rune, rune, error) {
	qNoise := rune(Clamp(math.Round(runesCount*fNoise), 0, runesCount))
	n, err := RandInt(2)

	if err != nil {
		return 0, 0, 0, 0, err
	}

	if n.Uint64() == 0 {
		return 0, utf8.MaxRune - qNoise, runesCount - qNoise, utf8.MaxRune, nil
	} else {
		return qNoise, utf8.MaxRune, 0, qNoise - 1, nil
	}
}

type key [4]runesMap

func GenKey(fNoise float64) (key, error) {
	var k key

	for i := range 4 {
		k[i] = make(runesMap)
	}

	// var arr key = []

}

func encrypt(in string, rMap runesMap, fMinNoise float32, fMaxNoise float32) (string, error) {
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
