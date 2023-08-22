package searching

func BinarySearch(list []int, target int) int {
	start := 0
	end := len(list) - 1
	for start <= end {
		middle := (start + end) / 2
		if list[middle] == target {
			return middle
		} else if list[middle] < target {
			start += 1
		} else {
			end -= 1
		}
	}

	return -1
}
