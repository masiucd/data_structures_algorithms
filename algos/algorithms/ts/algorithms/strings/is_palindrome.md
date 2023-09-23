# Is palindrome

## Description

Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

**Note:** For the purpose of this problem, we define empty string as valid palindrome.

## Complexity Analysis

- Time complexity: `O(n)` where `n` is the length of the string `s`.
- Space complexity: `O(1)` since we are not using any extra space to solve this problem in place.

### Example

```ts
isPalindrome("abba"); // true
isPalindrome("abbc"); // false
isPalindrome(""); // true
```

## Solution

_Using pointers_
Time complexity: O(n)
Space complexity: O(1)

```ts
function isPalindrome(input: string) {
  let start = 0;
  let end = input.length - 1;

  while (start < end) {
    if (input[start] === input[end]) {
      start++;
      end--;
    } else {
      return false;
    }
  }
  return true;
}
```

_Using recursion_
Time complexity: O(n)
Space complexity: O(n)

```ts
function isPalindrome(input: string) {
  if (input.length === 0 || input.length === 1) {
    return true;
  }

  if (input[0] === input[input.length - 1]) {
    return isPalindrome(input.slice(1, input.length - 1));
  }

  return false;
}
```
