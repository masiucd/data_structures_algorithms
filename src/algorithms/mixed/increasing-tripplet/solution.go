package mixed

import "math"

// increasingTriplet finds if there exists a triplet (i, j, k) such that nums[i] < nums[j] < nums[k].
// It returns true if such a triplet exists, otherwise it returns false.
// The function takes an input slice of integers, nums, and uses a two-pointer approach to find the triplet.
// It initializes two variables, a and b, with the maximum integer value.
// It then iterates through the input slice and updates the values of a and b accordingly.
// If a number is less than or equal to a, it updates the value of a.
// If a number is greater than a but less than or equal to b, it updates the value of b.
// If a number is greater than both a and b, it means a triplet exists and the function returns true.
// If no triplet is found after iterating through the entire slice, the function returns false.
func increasingTriplet(nums []int) bool {
	a, b := math.MaxInt32, math.MaxInt32
	for _, n := range nums {
		if n <= a {
			a = n
		} else if n <= b {
			b = n
		} else {
			return true
		}
	}
	return false
}
