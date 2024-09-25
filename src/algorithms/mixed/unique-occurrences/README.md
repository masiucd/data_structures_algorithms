# Unique Occurrences

Given an array of integers, determine if the number of occurrences of each value in the array is unique.

## Examples

- Input: [1, 2, 2, 1, 1, 3]
- Output: true
- Explanation: The number of occurrences of each value is unique (1 occurs 3 times, 2 occurs 2 times, and 3 occurs 1 time).

- Input: [1, 2]
- Output: false
- Explanation: The number of occurrences of each value is not unique (1 occurs 1 time and 2 occurs 1 time).

- Input: [-3, 0, 1, -3, 1, 1, 1, -3, 10, 0]
- Output: true
- Explanation: The number of occurrences of each value is unique (-3 occurs 3 times, 0 occurs 2 times, 1 occurs 4 times, and 10 occurs 1 time).

## Constraints

- 1 <= arr.length <= 1000
- -1000 <= arr[i] <= 1000

## Approach

- Use a map to count the occurrences of each value in the array.
- Use a set to check if the occurrences are unique.
- Return true if the occurrences are unique, false otherwise.

## Time Complexity

O(n)

## Space Complexity

O(n)

## Code

```go
func uniqueOccurrences(arr []int) bool {
	countMap := make(map[int]int)
	for _, num := range arr {
		countMap[num]++
	}

	occurrenceSet := make(map[int]bool)
	for _, count := range countMap {
		if occurrenceSet[count] {
			return false
		}
		occurrenceSet[count] = true
	}
	return true // All occurrence counts are unique
}
```

### Resources

- [LeetCode](https://leetcode.com/problems/unique-number-of-occurrences/)
- [GeeksforGeeks](https://www.geeksforgeeks.org/check-if-frequency-of-all-characters-are-same-in-a-string/)
