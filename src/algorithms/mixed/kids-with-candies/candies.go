package mixed

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var result []bool = make([]bool, len(candies))
	highestValue := maxValue(candies)
	for i, v := range candies {
		if v+extraCandies >= highestValue {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result

}

func maxValue(candies []int) int {
	if len(candies) == 0 {
		return 0
	}
	var highestValue int = candies[0]
	for i := 1; i < len(candies); i++ {
		if candies[i] > highestValue {
			highestValue = candies[i]
		}
	}
	return highestValue
}
