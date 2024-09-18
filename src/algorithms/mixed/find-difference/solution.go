package mixed

func findDifference(nums1, nums2 []int) [][]int {
	set1, set2 := make(map[int]bool), make(map[int]bool)
	var result [][]int = make([][]int, 2)

	for _, x := range nums1 {
		set1[x] = true
	}

	for _, x := range nums2 {
		set2[x] = true
	}

	var xs1 []int
	var xs2 []int
	for k := range set1 {
		if _, ok := set2[k]; !ok {
			xs1 = append(xs1, k)
		}
	}
	for k := range set2 {
		if _, ok := set1[k]; !ok {
			xs2 = append(xs2, k)
		}
	}
	result[0] = xs1
	result[1] = xs2
	return result
}
