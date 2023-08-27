# My own Math.pow

Creating my own Math.pow function.
Where the function receives two parameters, the base and the exponent.

## Example

```js
pow(2, 3); // 8
pow(2, 4); // 16
pow(2, 5); // 32
pow(2, 6); // 64
```

## Solution

```js
function pow(base, exponent) {
  let result = 1;
  for (let i = 0; i < exponent; i++) {
    result *= base;
  }
  return result;
}
```
