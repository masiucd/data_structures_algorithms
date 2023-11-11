# Sliding Window

## What is Sliding Window?

Sliding Window is a technique used to solve problems by using a window that slides over an array or string. The window is defined by two indices, the left and the right index. The window will increase or decrease in size depending on the problem and certain conditions. The technique is used to reduce the time complexity from O(n^2) to O(n).

## When to use Sliding Window?

Sliding Window is used to solve problems that require finding a subset of a contiguous subarray or substring that satisfies certain conditions. The technique is used to reduce the time complexity from O(n^2) to O(n).

## Example

Given an array of integers, find the maximum sum of a contiguous subarray of the array.

```go
func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currentSum := maxSum

    for i := 1; i < len(nums); i++ {
        currentSum = max(nums[i], currentSum + nums[i])
        maxSum = max(maxSum, currentSum)
    }

    return maxSum
}
```


## References

- [Wikipedia](https://en.wikipedia.org/wiki/Sliding_window_protocol)
- [GeeksforGeeks](https://www.geeksforgeeks.org/window-sliding-technique/)