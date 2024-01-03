package twosumii

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left <= right {
		tempTarget := numbers[left] + numbers[right]
		if tempTarget == target {
			return []int{left + 1, right + 1}
		} else if tempTarget < target {
			left++
		} else if tempTarget > target {
			right--
		}
	}
	return []int{0, 0}
}
