package sorting

func InsertionSort(list []int) []int {
	for i := 1; i < len(list); i++ {
		j := i - 1
		for j >= 0 && list[j+1] < list[j] {
			swap(j+1, j, list)
			j--
		}
	}
	return list
}
