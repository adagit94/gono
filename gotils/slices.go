package gotils

func MapSlice[S ~[]E, E any, EE any](s S, f func(e E, i int) EE) []EE {
	ss := make([]EE, len(s))

	for i, e := range s {
		ss[i] = f(e, i)
	}

	return ss
}