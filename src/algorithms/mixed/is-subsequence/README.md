# Is subsequence

Given a string s and a string t, check if s is subsequence of t.

A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (ie, "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:

```
Input: s = "abc", t = "ahbgdc"
Output: true
```

Example 2:

```
Input: s = "axc", t = "ahbgdc"
Output: false
```

Constraints:

- 0 <= s.length <= 100
- 0 <= t.length <= 10^4
- Both strings consists only of lowercase characters.

````

## Solution

```go
func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }
    if len(t) == 0 {
        return false
    }
    i, j := 0, 0
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i++
        }
        j++
    }
    return i == len(s)
}
````

### Complexity Analysis

The runtime complexity of this solution is O(n) where n is the length of the string t. The space complexity is O(1) since we are using only a constant amount of space.

- Time complexity: O(n)
- Space complexity: O(1)
- Where n is the length of the string t.
- The runtime complexity of this solution is O(n) where n is the length of the string t. The space complexity is O(1) since we are using only a constant amount of space.

### Resources

- [leetcode.com](https://leetcode.com/problems/is-subsequence/)
- [geeksforgeeks.org](https://www.geeksforgeeks.org/given-two-strings-find-first-string-subsequence-second/)
- [wikipedia.org](https://en.wikipedia.org/wiki/Subsequence)
