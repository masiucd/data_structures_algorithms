package sorting

// MergeSort algorithm
func MergeSort(nums []int) []int {
	if len(nums) < 2 { // base case
		return nums
	}
	middle := len(nums) / 2
	left := MergeSort(nums[:middle])
	right := MergeSort(nums[middle:])
	return merge(left, right)

}

func merge(a, b []int) []int {
	var final []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}

	return final
}
