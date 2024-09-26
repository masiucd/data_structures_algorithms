# Single Number

Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

## Examples

- Input: nums = [2,2,1]
- Output: 1

- Input: nums = [4,1,2,1,2]
- Output: 4

- Input: nums = [1]
- Output: 1

## Constraints

- 1 <= nums.length <= 3 \* 10^4
- -3 _ 10^4 <= nums[i] <= 3 _ 10^4
- Each element in the array appears twice except for one element which appears only once.

## Approach

**Using a map**

- Use a map to count the occurrences of each value in the array.
- Use a set to check if the occurrences are unique.
- Return true if the occurrences are unique, false otherwise.

**Using XOR**

- Use a variable to store the result of the XOR operation.
- Return the result.

## Time Complexity

O(n)

## Space Complexity

O(1)

## XOR

- XOR is a binary operation that returns 1 if the bits are different, otherwise 0.
- XOR is commutative and associative.
- XOR with 0 returns the number itself.
- XOR with itself returns 0.
- XOR is used in many algorithms, such as finding the single number.
