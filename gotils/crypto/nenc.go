package crypto

import (
	"fmt"
	"maps"
	"slices"
	n "github.com/adagit94/gono/gotils/numbers"
	sl "github.com/adagit94/gono/gotils/slices"
	uc "github.com/adagit94/gono/gotils/unicode"
)

type Replacements map[rune][]rune

type Key struct {
	Replacements   Replacements
	PotentialNoise []rune
}

func GenKey(reservedRunes []rune, replacementsCount int) Key {
	runes := make([]rune, uc.RunesCount)

	for r := range uc.RunesCount {
		runes[r] = r
	}

	replacementsRequired := len(reservedRunes) * replacementsCount

	if replacementsRequired > uc.RunesCount {
		panic(fmt.Errorf("Quantity of required runes cannot be higher than available count of runes (%v).", uc.RunesCount))
	}

	potentialReplacements := slices.Clone(runes)
	replacements := make(Replacements, len(reservedRunes))

	for range replacementsCount {
		for _, r := range reservedRunes {
			i := 0

			if l := int64(len(potentialReplacements)); l > 1 {
				i = int(n.RandIntPanic(l).Int64())
			}

			_, key := replacements[r]

			if !key {
				replacements[r] = make([]rune, 0)
			}

			replacements[r] = append(replacements[r], potentialReplacements[i])
			potentialReplacements = slices.Delete(potentialReplacements, i, i+1)
		}
	}

	potentialNoise := sl.Difference(runes, sl.Flat(slices.Collect(maps.Values(replacements))))

	return Key{Replacements: replacements, PotentialNoise: potentialNoise}
}

func Encrypt(in string, key Key, noiseMin int64, noiseMax int64) string {
	out := make([]rune, len(in))

	for i, r := range in {
		replacements, replacementsKey := key.Replacements[r]

		if !replacementsKey {
			panic(fmt.Errorf("No replacement available for character %v. Rune key isn't present.", r))
		}

		out[i] = replacements[n.RandIntPanic(int64(len(replacements))).Int64()]
	}

	if qPotentialNoise, qNoise := int64(len(key.PotentialNoise)), noiseMin+n.RandIntPanic(noiseMax-noiseMin+1).Int64(); qPotentialNoise > 0 && qNoise > 0 {
		for range qNoise {
			i := n.RandIntPanic(qPotentialNoise).Int64()
			j := 0

			if l := len(out); l > 0 {
				j = int(n.RandIntPanic(int64(len(out) + 1)).Int64())
			}

			out = slices.Insert(out, j, key.PotentialNoise[i])
		}
	}

	return string(out)
}

func Decrypt(in string, key Key) string {
	out := []rune{}
	underlayingRunes := maps.Keys(key.Replacements)

	for _, r := range in {
		for underlayingRune := range underlayingRunes {
			rpmMatched := slices.Contains(key.Replacements[underlayingRune], r)

			if rpmMatched {
				out = append(out, underlayingRune)
				break
			}
		}
	}

	return string(out)
}
