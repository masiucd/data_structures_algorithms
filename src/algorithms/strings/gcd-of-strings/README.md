# GCD of strings

Given two strings `str1` and `str2` of sizes `n` and `m` respectively, the GCD string of the two strings is the largest string that divides both `str1` and `str2`.

The GCD string can be determined by finding the greatest common divisor of the lengths of the two strings and then checking if the string is a common divisor of both strings.

## Example 1:

```
Input: str1 = "ABCABC", str2 = "ABC"
Output: "ABC"
Explanation: The GCD string is "ABC" which divides both strings.
```

## Example 2:

```
Input: str1 = "ABABAB", str2 = "ABAB"
Output: "AB"
Explanation: The GCD string is "AB" which divides both strings.
```

## Example 3:

```

Input: str1 = "LEET", str2 = "CODE"
Output: ""
Explanation: There is no common divisor of the two strings.
```

## Constraints:

- `1 <= str1.length, str2.length <= 1000`
- `str1` and `str2` consist of uppercase and lowercase English letters.
- `str1` and `str2` contain at most 1000 unique characters.
- `str1` and `str2` contain only characters from the English alphabet.
- `str1` and `str2` are non-empty strings.
- `str1` and `str2` are not equal.

## Hints

1. The GCD string is a common divisor of both strings.
2. The GCD string is a prefix of both strings.
3. The GCD string is a suffix of both strings.
4. The GCD string is a substring of both strings.
5. The GCD string is a common divisor of both strings and is the longest string that divides both strings.
6. The GCD string is a common divisor of both strings and is the longest string that divides both strings. The GCD string is unique.

### Reference

- [LeetCode](https://leetcode.com/problems/greatest-common-divisor-of-strings/)
- [GeeksForGeeks](https://www.geeksforgeeks.org/gcd-of-two-strings/)
