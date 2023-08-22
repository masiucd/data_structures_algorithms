package searching

type numOrStr interface {
	int | string
}

func LinearSearch[T numOrStr](list []T, target T) int {
	for i, v := range list {
		if v == target {
			return i
		}
	}
	return -1
}
