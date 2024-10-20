# Consecutive Numbers

## Level 1

In Level 1, you are tasked with writing a function that takes a non-empty list of integers and checks if they can be rearranged to form a consecutive sequence of numbers. The function should return the sorted list if the list is consecutive, or an empty list if it's not.

### Example

```Go
ConsecutiveNumbers([]int{1,3}) // []int{}
// because there is no way to rearrange these numbers to form a consecutive sequence.

ConsecutiveNumbers([]int{1, 0, -1}) // []int{-1, 0, 1}
// because these numbers can be rearranged to form a consecutive sequence.
```

## Level 2

In Level 2, you are tasked with writing a function that takes a non-empty list of integers and returns all the integers missing to complete a list of consecutive integers (as described in Level 1). The list should be sorted in ascending order. If no integer is missing, return an empty list.

### Example

```Go
ConsecutiveNumbers([]int{1, 3, 4, 7, 9}) // []int{2, 5, 6, 8}
// because the missing numbers are 2, 5, 6, and 8.
```

## Hints

1. You can use the `sort` package to sort the list of integers.
2. You can use the `math` package to get the absolute value of a number.
