package binarysearch

func binarySearch(xs []int, target int) int {
	start := 0
	end := len(xs) - 1
	for start <= end {
		middle := (start + end) / 2
		if xs[middle] == target {
			return middle
		}
		if xs[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}
