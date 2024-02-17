# Palindrome rearranging

## Description

Given a string, find out if its characters can be rearranged to form a palindrome.

## Example

For `inputString = "aabb"`, the output should be `palindromeRearranging(inputString) = true`.

We can rearrange `"aabb"` to make `"abba"`, which is a palindrome.

For `inputString = "abca"`, the output should be `palindromeRearranging(inputString) = false`.

## Time and Space Complexity

The time complexity is `O(n)` and the space complexity is `O(n)`, where `n` is the length of the input string.

## Solution

The solution is to count the number of characters in the string and check if the number of characters is even or odd. If the number of characters is even, then all characters should have an even count. If the number of characters is odd, then only one character should have an odd count and all other characters should have an even count.

The solution is implemented in the `palindromeRearranging` method.

```go
func palindromeRearranging(inputString string) bool {
    charCount := make(map[rune]int)
    for _, char := range inputString {
        charCount[char]++
    }
    oddCount := 0
    for _, count := range charCount {
        if count % 2 != 0 {
            oddCount++
        }
    }
    return oddCount <= 1
}
```
