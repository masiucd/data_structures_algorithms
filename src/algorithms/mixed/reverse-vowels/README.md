# Reverse vowels

Given a string, reverse only the vowels of the string.

Example 1:

```
Input: "hello"
Output: "holle"
```

Example 2:

```
Input: "leetcode"
Output: "leotcede"
```

## Solution

The solution is simple. We can use two pointers to traverse the string from both ends. If both characters are vowels, we swap them. If one of them is not a vowel, we move the pointer to the next character. We continue this process until the two pointers meet.

The time complexity of this solution is O(n), where n is the length of the string. The space complexity is O(1).

```go
package main

import "fmt"

func reverseVowels(s string) string {
  vowels := map[byte]bool{
    'a': true,
    'e': true,
    'i': true,
    'o': true,
    'u': true,
    'A': true,
    'E': true,
    'I': true,
    'O': true,
    'U': true,
  }

  n := len(s)
  sBytes := []byte(s)
  i, j := 0, n-1

  for i < j {
    for i < j && !vowels[sBytes[i]] {
      i++
    }
    for i < j && !vowels[sBytes[j]] {
      j--
    }

    sBytes[i], sBytes[j] = sBytes[j], sBytes[i]
    i++
    j--
  }

  return string(sBytes)
}
```

**Complexity Analysis**

The time complexity of this solution is O(n), where n is the length of the string.

The space complexity is O(1).

### Resources

- [LeetCode](https://leetcode.com/problems/reverse-vowels-of-a-string/)
- [ProgramCreek](https://www.programcreek.com/2015/04/leetcode-reverse-vowels-of-a-string-java/)
