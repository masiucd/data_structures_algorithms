package mixed

// maxArea calculates the maximum area of water that can be contained
// between two lines represented by the heights in the input slice.
// The function uses a two-pointer approach to find the maximum area
// by iterating from both ends of the slice towards the center.
//
// Parameters:
// - height: A slice of integers representing the heights of the lines.
//
// Returns:
// - An integer representing the maximum area of water that can be contained.
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	area := 0
	for left < right {
		minHeight := min(height[left], height[right])
		width := right - left
		area = max(area, minHeight*width)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return area
}
