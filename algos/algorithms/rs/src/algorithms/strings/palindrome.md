# is palindrome

is_palindrome checks if a string is a palindrome.
for example "abba" is a palindrome, but "abab" is not.

## Example

```rs
assert_eq!(is_palindrome("abba"), true);
assert_eq!(is_palindrome("abab"), false);
```

## Complexity Analysis

- Time complexity: `O(n)` where `n` is the length of the string `s`.
- Space complexity: `O(1)` since we are not using any extra space to solve this problem in place.

## References

[Wikipedia](https://en.wikipedia.org/wiki/Palindrome)
[GeeksForGeeks](https://www.geeksforgeeks.org/c-program-check-given-string-palindrome/)
