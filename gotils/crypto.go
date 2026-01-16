package gotils

import (
	"fmt"
	"maps"
	"math"
	"slices"
	"unicode/utf8"
)

const runesCount = utf8.MaxRune + 1
const (
	n0, n9   uint8 = 48, 57
	aUp, zUp uint8 = 65, 90
	aLo, zLo uint8 = 97, 122
)

type runesMap map[rune][]rune

func genRunesSets(reservedRunes []rune, replacementsCount int) (runesMap, []rune) {
	runes := make([]rune, runesCount)

	for r := range runesCount {
		runes[r] = r
	}

	replacementsRequired := len(reservedRunes) * replacementsCount

	if replacementsRequired > runesCount {
		panic(fmt.Errorf("Quantity of required runes cannot be higher than available count of runes (%w).", runesCount))
	}

	potentialReplacements := slices.Clone(runes)
	replacements := make(runesMap, len(reservedRunes))

	for range replacementsCount {
		for _, r := range reservedRunes {
			i := 0

			if l := int64(len(potentialReplacements)); l > 1 {
				i = int(RandIntPanic(l).Int64())
			}

			_, key := replacements[r]

			if !key {
				replacements[r] = make([]rune, 0)
			}

			replacements[r] = append(replacements[r], potentialReplacements[i])
			potentialReplacements = slices.Delete(potentialReplacements, i, i+1)
		}
	}

	potentialNoise := Difference(runes, Flat(slices.Collect(maps.Values(replacements))))

	return replacements, potentialNoise
}

type key struct {
	replacements   runesMap
	potentialNoise []rune
}

func GenKey(reservedRunes []rune, replacementsCount int) key {
	replacements, potentialNoise := genRunesSets(reservedRunes, replacementsCount)
	k := key{replacements: replacements, potentialNoise: potentialNoise}

	return k
}

func encrypt(in string, key key, qNoiseMin int64, qNoiseMax int64) string {
	out := make([]rune, len(in))

	for i, r := range in {
		replacements, replacementsKey := key.replacements[r]

		if !replacementsKey {
			panic(fmt.Errorf("No replacement available for character %w. Rune key isn't present.", r))
		}

		out[i] = replacements[RandIntPanic(int64(len(replacements))).Int64()]
	}

	if qPotentialNoise, qNoise := int64(len(key.potentialNoise)), qNoiseMin+RandIntPanic(qNoiseMax-qNoiseMin+1).Int64(); qPotentialNoise > 0 && qNoise > 0 {
		for range qNoise {
			i := RandIntPanic(qPotentialNoise).Int64()
			j := int(RandIntPanic(int64(len(out))).Int64())

			out = slices.Insert(out, j, key.potentialNoise[i])
		}
	}


}
