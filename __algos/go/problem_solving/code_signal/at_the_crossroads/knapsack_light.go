package at_the_crossroads

func knapSackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
	if weight1+weight2 <= maxW {
		return value1 + value2
	} else if weight1 <= maxW && weight2 > maxW {
		return value1
	} else if weight2 <= maxW && weight1 > maxW {
		return value2
	} else if weight1 <= maxW && weight2 <= maxW {
		return max(value1, value2)
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func knapSackLight2(value1, weight1, value2, weight2, maxW int) int {
	switch {
	case weight1+weight2 <= maxW:
		return value1 + value2
	case weight1 <= maxW && weight2 > maxW:
		return value1
	case weight2 <= maxW && weight1 > maxW:
		return value2
	case weight1 <= maxW && weight2 <= maxW:
		if value1 > value2 {
			return value1
		} else {
			return value2
		}
	default:
		return 0
	}
}
