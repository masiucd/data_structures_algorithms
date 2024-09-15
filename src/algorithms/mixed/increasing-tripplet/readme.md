# Increasing triplet subsequence

Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that i < j < k and `nums[i]` < `nums[j]` < `nums[k]`. If no such indices exist, return false.

Example 1:

```
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.
```

Example 2:

```
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
```

Example 3:

```
Input: nums = [2,1,5,0,4,6]
Output: true
Explanation: The triplet (2, 5, 6) is valid because nums[0] < nums[2] < nums[5].
```

### Description

Initialization:

Two variables, first and second, are initialized to the maximum possible integer value (math.MaxInt32). These variables will keep track of the smallest and second smallest elements encountered so far in the slice.
Iteration:

The function iterates through each element n in the input slice nums.
Updating first and second:

If the current element n is smaller than first, it updates first to n. This means n is the smallest element encountered so far.
If n is not smaller than first but is smaller than second, it updates second to n. This means n is the second smallest element encountered so far, but larger than first.
If n is greater than both first and second, it means a triplet has been found (first < second < n), and the function returns true.
Completion:

If the loop completes without finding such a triplet, the function returns false.

### Complexity Analysis

The time complexity for this approach is O(n), where n is the length of the input slice nums. The function iterates through each element in the slice once.

The space complexity is O(1) since the function uses only a constant amount of extra space for the two variables first and second.

### Resources

- [LeetCode Official Problem Page](https://leetcode.com/problems/increasing-triplet-subsequence/)
- [GeeksforGeeks Article on Increasing Triplet Subsequence](https://www.geeksforgeeks.org/find-a-sorted-subsequence-of-size-3-in-linear-time/)
