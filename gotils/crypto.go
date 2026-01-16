package gotils

import (
	"fmt"
	"maps"
	"slices"
	"unicode/utf8"
)

const runesCount = utf8.MaxRune + 1
const (
	n0, n9   uint8 = 48, 57
	aUp, zUp uint8 = 65, 90
	aLo, zLo uint8 = 97, 122
)

type runesMap map[rune]rune

func genRunesSets(reservedRunes []rune) (runesMap, []rune) {
	runes := make([]rune, runesCount)

	for r := range runesCount {
		runes[r] = r
	}

	replacementsRequired := len(reservedRunes)

	if replacementsRequired > runesCount {
		panic(fmt.Errorf("Quantity of reserved runes cannot be higher than available count of runes (%w).", runesCount))
	}

	potentialReplacements := slices.Clone(runes)
	replacements := make(runesMap, replacementsRequired)

	for _, r := range reservedRunes {
		i := 0

		if l := len(potentialReplacements); l > 1 {
			n, err := RandInt(int64(l))

			if err != nil {
				panic(err)
			}

			i = int(n.Uint64())
		}

		replacements[r] = potentialReplacements[i]
		potentialReplacements = slices.Delete(potentialReplacements, i, i+1)
	}

	potentialNoise := Difference(runes, slices.Collect(maps.Values(replacements)))

	return replacements, potentialNoise
}

type key struct {
	replacements   runesMap
	potentialNoise []rune
}

func GenKey(reservedRunes []rune) key {
	replacements, potentialNoise := genRunesSets(reservedRunes)
	k := key{replacements: replacements, potentialNoise: potentialNoise}

	return k
}

func encrypt(in string, key key, fMinNoise float64, fMaxNoise float64) string {
	out := make([]rune, len(in))

	for i, r := range in {
		replacement, k := key.replacements[r]

		if !k {
			panic(fmt.Errorf("No replacement available for character %w. Rune key isn't present.", r))
		}
		
		out[i] = replacement
	}

	noiseAvailableRunes := Difference(slices.Sorted(maps.Values(m)), noiseUnavailableRunes)
	qNoiseAvailableRunes := float64(len(noiseAvailableRunes))
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
