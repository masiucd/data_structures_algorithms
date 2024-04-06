package maxsum

func maxSum(heights []int) int {
	res := 0
	l, r := 0, len(heights)-1
	for l < r {
		width, height := r-l, min(heights[l], heights[r])
		area := width * height
		res = max(area, res)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
