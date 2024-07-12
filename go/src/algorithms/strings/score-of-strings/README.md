# Score of strings

You are given a string s. The score of a string is defined as the sum of the absolute difference between the ASCII values of adjacent characters.
Return the score of s.

## Example

```go

fmt.Println(ScoreOfStrings("zaz")) // 50
fmt.Println(ScoreOfStrings("hello")) // 13

```

### Example 1

Input: s = "zaz"
Output: 50
Explanation: The score of the string is |'z' - 'a'| + |'a' - 'z'| = 25 + 25 = 50.

### Example 2

Input: s = "hello"
Output: 13
Explanation: The score of the string is |'h' - 'e'| + |'e' - 'l'| + |'l' - 'l'| + |'l' - 'o'| = 1 + 2 + 0 + 10 = 13.

## Constraints

- 1 <= s.length <= 100
- s consists of lowercase English letters.
