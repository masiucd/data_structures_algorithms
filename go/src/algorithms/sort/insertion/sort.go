package insertion

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		// as long j is in bounce and arr[j] < arr[j-1], then we swap
		// current is less than previous
		for j > 0 && arr[j] < arr[j-1] {
			//swapping
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}
	return arr
}
