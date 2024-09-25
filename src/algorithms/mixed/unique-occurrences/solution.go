package mixed

// uniqueOccurrences checks if the number of occurrences of each value in the array is unique.
// It returns true if the occurrence count for every value is unique, and false otherwise.
// This function uses a more efficient approach by using a single loop to count occurrences
// and check for uniqueness simultaneously.
func uniqueOccurrences(arr []int) bool {
	intCount := make(map[int]int)
	for _, n := range arr {
		intCount[n]++
	}

	set := make(map[int]bool)
	for _, v := range intCount {
		if _, ok := set[v]; !ok {
			set[v] = true
		} else {
			return false
		}
	}
	return true
}

// uniqueOccurrences2 checks if the number of occurrences of each value in the array is unique.
// It returns true if the occurrence count for every value is unique, and false otherwise.
// This function uses a more efficient approach by using a single loop to count occurrences
// and check for uniqueness simultaneously.
func uniqueOccurrences2(arr []int) bool {

	// Create a map to count occurrences of each number
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	// Create a set to check for unique occurrence counts
	occurrenceSet := make(map[int]bool)
	for _, count := range countMap {
		if occurrenceSet[count] {
			return false // If count already exists, it's not unique
		}
		occurrenceSet[count] = true
	}

	return true // All occurrence counts are unique

}
