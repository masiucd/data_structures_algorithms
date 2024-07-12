# Array Maximal Adjacent Difference

## Description

Given an array of integers, find the maximal absolute difference between any two of its adjacent elements.

## Example

For `inputArray = [2, 4, 1, 0]`, the output should be `3`.

## Time Complexity

The time complexity is `O(n)`.

## Space Complexity

The space complexity is `O(1)`.

## Solution

The solution is to iterate through the array and find the maximal absolute difference between any two of its adjacent elements.

```go
func arrayMaximalAdjacentDifference(inputArray []int) int {
  maxDiff := 0
  for i := 1; i < len(inputArray); i++ {
    diff := int(math.Abs(float64(inputArray[i] - inputArray[i-1])))
    if diff > maxDiff {
      maxDiff = diff
    }
  }
  return maxDiff
}
```
