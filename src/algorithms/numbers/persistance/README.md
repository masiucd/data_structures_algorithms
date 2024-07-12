# Persistence

In mathematics, the persistence of a number is the number of times you have to replace the number with the sum of its digits until you reach a single digit.

Write a function that returns the persistence of a number.

## Example

```go
fmt.Println(Persistence(39)) // 3
fmt.Println(Persistence(999)) // 4
fmt.Println(Persistence(4)) // 0
```

// 39 --> 3 (because 3*9 = 27, 2*7 = 14, 1*4 = 4 and 4 has only one digit, there are 3 multiplications)
// 999 --> 4 (because 9*9*9 = 729, 7*2*9 = 126, 1*2*6 = 12, and finally 1*2 = 2, there are 4 multiplications)
// 4 --> 0 (because 4 is already a one-digit number, there is no multiplication)
