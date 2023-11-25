# FizzBuzz

Write a program that outputs the string/list-of-strings representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

## Example

```text
n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
```

## Solution

### Go

```go
func fizzBuzz(n int) []string {
    var result []string
    for i := 1; i <= n; i++ {
        if i % 15 == 0 {
            result = append(result, "FizzBuzz")
        } else if i % 3 == 0 {
            result = append(result, "Fizz")
        } else if i % 5 == 0 {
            result = append(result, "Buzz")
        } else {
            result = append(result, strconv.Itoa(i))
        }
    }
    return result
}
```
