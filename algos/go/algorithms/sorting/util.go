package sorting

func swap(i, j int, list []int) {
	a, b := list[i], list[j]
	list[i], list[j] = b, a
}
