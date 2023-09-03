# Count vowels

## Problem

Given a string, count the number of vowels in it and return that count.

### Example

```ts
countVowels("hello"); // 2
countVowels("why"); // 0
countVowels("aeiou"); // 5
```

## Solution

**iterative solution**

- Time Complexity: O(n)
- Space Complexity: O(1)

```ts
function countVowels(str: string): number {
  let vowels = ["a", "e", "i", "o", "u"];
  let count = 0;
  for (let i = 0; i < str.length; i++) {
    if (vowels.includes(str[i])) {
      count++;
    }
  }
  return count;
}
```

**recursive solution**

- Time Complexity: O(n)
- Space Complexity: O(n)

```ts
function countVowels(str: string): number {
  let vowels = ["a", "e", "i", "o", "u"];
  if (str.length === 0) {
    return 0;
  }
  if (vowels.includes(str[0])) {
    return 1 + countVowels(str.slice(1));
  } else {
    return countVowels(str.slice(1));
  }
}
```
