package assert

func Assert[T any](x any) T {
	return x.(T)
}

func AssertSafely[T any](x any) (T, bool) {
	v, ok := x.(T)
	return v, ok
}