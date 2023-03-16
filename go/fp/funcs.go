package fp

func MapList[T any](xs []T, fn func(T) T) []T {
	var result []T
	for _, v := range xs {
		result = append(result, fn(v))
	}
	return result
}

func FilterList[T any](xs []T, fn func(T) bool) []T {
	var result []T
	for _, v := range xs {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}
