package mixed

import "math"

func arrayMaximalAdjacentDifference(inputList []int) int {
	var max int
	for i := 1; i < len(inputList); i++ {
		prev := inputList[i-1]
		curr := inputList[i]
		diff := maxDifference(prev, curr)
		if max < diff {
			max = diff
		}
	}
	return max
}

func maxDifference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func arrayMaximalAdjacentDifferenceV2(inputArray []int) int {
	maxDiff := 0
	for i := 1; i < len(inputArray); i++ {
		diff := int(math.Abs(float64(inputArray[i] - inputArray[i-1])))
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}
