# Reverse list using stack

Given a list of integers, reverse it using stack.

## Example

**Input** : `[1, 2, 3, 4, 5]`

**Output**: `[5, 4, 3, 2, 1]`

## Solution

```ts
/**
 * Reverse list using stack
 * @param list - List of integers
 * @returns Reversed list
 */

export function reverse(list: number[]): number[] {
  const stack: number[] = [];
  const reversedList: number[] = [];

  for (const item of list) {
    stack.push(item);
  }

  while (stack.length > 0) {
    reversedList.push(stack.pop()!);
  }

  return reversedList;

  // Alternatively, you can use Array.reverse() method
```
