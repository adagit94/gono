package maps

import (
	"maps"
	"slices"
)

func Swap[K comparable, V comparable](m map[K]V) map[V]K {
	keys := slices.Collect(maps.Keys(m))
	m2 := make(map[V]K, len(keys))

	for _, k := range keys {
		m2[m[k]] = k
	}
	
	return m2
}
