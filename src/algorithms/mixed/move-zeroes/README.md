# Move Zeroes

Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

**Note** that you must do this in-place without making a copy of the array.

**Example 1:**

```
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
```

**Example 2:**

```

Input: nums = [0]
Output: [0]
```

**Constraints:**

- `1 <= nums.length <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`
- You must do this in-place without making a copy of the array.
- Minimize the total number of operations.

````

## Solution

The solution is to use two pointers, one to iterate the array and another to keep track of the position where the next non-zero element should be placed.

1. Iterate the array with the pointer `i` from `0` to `n-1`.
2. If the element at the current position is not `0`, then swap the element at the current position with the element at the position of the pointer `j`.
3. Increment the pointer `j` by `1`.
4. Continue the process until the pointer `i` reaches the end of the array.
5. Fill the remaining positions from the pointer `j` to the end of the array with `0`.

The time complexity of this solution is `O(n)`.

```go

func moveZeroes(nums []int) {
    n := len(nums)
    j := 0
    for i := 0; i < n; i++ {
        if nums[i] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            j++
        }
    }
    for j < n {
        nums[j] = 0
        j++
    }
}
````

### References

- [LeetCode](https://leetcode.com/problems/move-zeroes/)
- [ProgramCreek](https://www.programcreek.com/2014/05/leetcode-move-zeroes-java/)
- [GeeksforGeeks](https://www.geeksforgeeks.org/move-zeroes-end-array/)
