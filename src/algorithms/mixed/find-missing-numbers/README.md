# Find missing numbers

## Description

Given an array of integers where each value is between 1 and N (N is the length of the array), find all the duplicates in the array.

## Example

```go
fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})) // [5, 6]
```

```go
fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1, 5, 5})) // [6, 9, 10]
```

```go
fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1, 5, 6})) // []
```

## Constraints

- The array should not be modified
- The result should be sorted
- The result should not contain duplicates
- The result should be in the range of 1 to N
