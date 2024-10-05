# Number game

You are given a 0-indexed integer array nums of even length and there is also an empty array arr. Alice and Bob decided to play a game where in every round Alice and Bob will do one move. The rules of the game are as follows:

Every round, first Alice will remove the minimum element from nums, and then Bob does the same.
Now, first Bob will append the removed element in the array arr, and then Alice does the same.
The game continues until nums becomes empty.
Return the resulting array arr.

## Examples

```
Input: nums = [1,3,4,2]
Output: [2,4,1,3]
```

```
Input: nums = [1,2]
Output: [2,1]
```

```
Input: nums = [2,2,2,2]
Output: [2,2,2,2]
```

## Constraints

- 2 <= nums.length <= 2 \* 10^5
- 1 <= nums[i] <= 10^5
- nums.length is even.

## Approach

- Sort the array
- Swap adjacent elements
- Return the result

## Code

```go
func numberGame(nums []int) []int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}
```
