# Data structures and algorithms Golang

## About
Repository containing data structures and algorithms implemented using the Go programming language.


### Content

#### Data structures

* Single linked list
* Double linked list
* Stack
* Queue
* Binary search tree

#### Algorithms



### FP 
Some handy FP functions using generics

```go
  func MapList[T any](xs []T, fn func(T) T) []T {
    var result []T
    for _, v := range xs {
      result = append(result, fn(v))
    }
    return result
  }

  func FilterList[T any](xs []T, fn func(T) bool) []T {
    var result []T
    for _, v := range xs {
      if fn(v) {
        result = append(result, v)
      }
    }
    return result
  }

    nums := []int{1, 2, 3, 4}
    xs := fp.MapList(nums, func(n int) int { return n * 2 })
    fmt.Println(xs) //[2 4 6 8]


    xs = fp.FilterList(xs, func(n int) bool { return n > 4 })
    fmt.Println(xs) // [6 8]
```