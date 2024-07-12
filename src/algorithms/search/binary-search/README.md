# Binary search

## About

Binary search is a search algorithm that finds the position of a target value within a sorted array.

Binary search compares the target value to the middle element of the array. If they are not equal, the half in which the target cannot lie is eliminated and the search continues on the remaining half, again taking the middle element to compare to the target value, and repeating this until the target value is found. If the search ends with the remaining half being empty, the target is not in the array.

## Algorithm

Given an array `arr` of `n` elements with values or records `arr[0], arr[1], ..., arr[n-1]`, sorted such that `arr[0] <= arr[1] <= ... <= arr[n-1]`, and target value `target`, the following subroutine uses binary search to find the index of `target` in `arr`.

```go
func binarySearch(xs []int, target int) int {
	start := 0
	end := len(xs) - 1
	for start <= end {
		middle := (start + end) / 2
		if xs[middle] == target {
			return middle
		}
		if xs[middle] < target {
			start = middle + 1
		} else {
			end = middle - 1
		}
	}
	return -1
}

```

## Complexity

| Name              | Best |  Average  |   Worst   | Memory | Stable |
| ----------------- | :--: | :-------: | :-------: | :----: | :----: |
| **Binary search** | Θ(1) | Θ(log(n)) | Θ(log(n)) |  Θ(1)  |  Yes   |

## References

[Wikipedia](https://en.wikipedia.org/wiki/Binary_search_algorithm)
