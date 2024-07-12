package fp

// Filter returns a function that filters a slice of integers based on a predicate function.
// The predicate function defines a condition that each element in the slice must satisfy to be included in the result.
//
// Parameters:
//   - predicate: A function that takes an integer as input and returns a boolean.
//     This function is used to test each element of the input slice. If the function returns true for an element,
//     that element is included in the result slice.
//
// Returns:
//   - A function that takes a slice of integers (xs []int) as input and returns a new slice of integers.
//     This returned function iterates over each element in the input slice (xs) and applies the predicate function to it.
//     If the predicate function returns true for an element, that element is added to the result slice.
//     The function finally returns the result slice containing all elements that satisfy the predicate.
//
// Example:
// isEven := func(n int) bool { return n%2 == 0 }
// filterEven := Filter(isEven)
// evens := filterEven([]int{1, 2, 3, 4, 5})
// fmt.Println(evens) // Output: [2, 4]
func Filter(predicate func(n int) bool) func(xs []int) []int {
	return func(xs []int) []int {
		var result []int
		for _, x := range xs {
			if predicate(x) {
				result = append(result, x)
			}
		}
		return result
	}
}
