# Product except self

Given an array `nums` of `n` integers where `n > 1`, return an array `output` such that `output[i]` is equal to the product of all the elements of `nums` except `nums[i]`.

**Example:**

```
Input:  [1,2,3,4]
Output: [24,12,8,6]
```

**Note:** Please solve it without division and in O(n).

## Solution

The solution is to calculate the product of all elements to the left of the current element and then multiply it by the product of all elements to the right of the current element.

The algorithm is as follows:

1. Create a list `output` of the same size as the input list `nums` and initialize it with 1.
2. Initialize two variables `left` and `right` to 1.
3. Iterate over the input list `nums` from left to right.
4. Update the current element of the `output` list by multiplying it with the `left` variable.
5. Update the `left` variable by multiplying it with the current element of the input list `nums`.
6. Repeat the above steps until the end of the input list `nums`.
7. Iterate over the input list `nums` from right to left.
8. Update the current element of the `output` list by multiplying it with the `right` variable.
9. Update the `right` variable by multiplying it with the current element of the input list `nums`.
10. Repeat the above steps until the start of the input list `nums`.
11. Return the `output` list as the final result.
12. The time complexity of this algorithm is O(n), where n is the size of the input list `nums`.
13. The space complexity of this algorithm is O(n), where n is the size of the input list `nums`.
14. The code implementation of this algorithm is as follows:

```go
package main

import "fmt"

func productExceptSelf(nums []int) []int {
    n := len(nums)
    output := make([]int, n)
    left, right := 1, 1

    for i := 0; i < n; i++ {
        output[i] = left
        left *= nums[i]
    }

    for i := n - 1; i >= 0; i-- {
        output[i] *= right
        right *= nums[i]
    }

    return output
}



func main() {
    nums := []int{1, 2, 3, 4}
    fmt.Println(productExceptSelf(nums)) // Output: [24 12 8 6]
}
```

### References

- [LeetCode Problem](https://leetcode.com/problems/product-of-array-except-self/)
- [ProgramCreek](https://www.programcreek.com/2014/07/leetcode-product-of-array-except-self-java/)
