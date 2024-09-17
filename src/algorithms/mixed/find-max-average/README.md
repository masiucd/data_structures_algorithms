# Find Max Average

Given an array consisting of `n` integers, find the contiguous subarray of given length `k` that has the maximum average value. And you need to output the maximum average value.

## Example 1

```text
Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75

Explanation: Maximum average is (12-5-6+50)/4 = 51/4 = 12.75
```

## Note

1. 1 <= `k` <= `n` <= 30,000.
2. Elements of the given array will be in the range `[-10,000, 10,000]`.

## Hints

1. Sliding window.
2. What is the time complexity of the sliding window approach?
3. How will you optimize the sliding window approach?

### Mental model for sliding window

The sliding window technique is a way to increase the efficiency of a brute force algorithm. It is a way to solve problems by only considering a subset of the elements in the input array. The subset is usually a contiguous subarray. The window is the subset of elements that you are currently considering. You slide the window by removing the first element of the window and adding a new element to the end of the window.

The sliding window technique is useful when you have to find a substring, subarray, or a desired value that is continuous in the input array. The technique is also useful when you have to find the longest substring, subarray, or the smallest or largest value that satisfies a certain condition.

The sliding window technique is a way to solve problems by only considering a subset of the elements in the input array. The subset is usually a contiguous subarray. The window is the subset of elements that you are currently considering. You slide the window by removing the first element of the window and adding a new element to the end of the window.

1. Initialize the window with the first `k` elements of the array.
2. Compute the sum of the elements in the window.
3. Slide the window by removing the first element of the window and adding a new element to the end of the window.
4. Compute the sum of the elements in the window.
5. Repeat steps 3 and 4 until the window reaches the end of the array.
6. Keep track of the maximum sum of the elements in the window.
7. Return the maximum sum of the elements in the window.
8. The maximum average is the maximum sum divided by `k`.
9. Return the maximum average.
10. The time complexity of the sliding window approach is `O(n)`.
11. The space complexity of the sliding window approach is `O(1)`.
12. The sliding window technique is a way to increase the efficiency of a brute force algorithm. It is a way to solve problems by only considering a subset of the elements in the input array. The subset is usually a contiguous subarray. The window is the subset of elements that you are currently considering. You slide the window by removing the first element of the window and adding a new element to the end of the window.
