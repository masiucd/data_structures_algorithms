package alllongeststrings

func allLongestStrings(inputArray []string) []string {
	var result []string
	longest := 0
	for _, s := range inputArray {
		if len(s) > longest {
			longest = len(s)
			result = []string{s}
		} else if len(s) == longest {
			result = append(result, s)
		}
	}
	return result
}
