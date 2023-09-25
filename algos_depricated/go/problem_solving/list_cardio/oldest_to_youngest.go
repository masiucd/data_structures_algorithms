package list_cardiopackage

import "sort"

//Sort the inventors by birthdate, oldest to youngest

func oldestToYoungest(input []Inventor) []Inventor {
	sort.SliceStable(input, func(a, b int) bool {
		return input[a].year < input[b].year
	})
	return input
}
