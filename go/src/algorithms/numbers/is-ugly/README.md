# Is ugly number

An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

## Solution

```go
package main

import "fmt"

func isUgly(num int) bool {
  if num <= 0 {
    return false
  }
  for _, i := range []int{2, 3, 5} {
    for num%i == 0 {
      num /= i
    }
  }
  return num == 1
}
```

The `isUgly()` function takes an integer as input and returns a boolean value. It checks if the given number is an ugly number or not. An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5. The function first checks if the input number is less than or equal to 0. If it is, the function returns false. The function then iterates over the numbers 2, 3, and 5. For each number, it divides the input number by that number as long as the input number is divisible by that number. If the input number is divisible by all three numbers, the function returns true; otherwise, it returns false.
