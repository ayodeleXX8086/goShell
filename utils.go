package main

func find[T any](arr []T, pred func(T) bool) (T, bool) {
	for _, v := range arr {
		if pred(v) {
			return v, true
		}
	}
	var defaultT T
	return defaultT, false
}
