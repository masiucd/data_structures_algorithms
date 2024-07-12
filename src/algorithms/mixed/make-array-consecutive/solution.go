package makearrayconsecutive

import "go-ds/src/algorithms/sort/quick"

func makeArrayConsecutive(statues []int) (result int) {
	// sort the slice
	xs := quick.QuickSort(statues)
	for i := 0; i < len(xs)-1; i++ {
		current := xs[i]
		next := xs[i+1]
		// get the difference between the next and the current value
		diff := next - current
		// if the diff is greater then one then add it to the result
		if diff > 1 {
			result += diff - 1
		}
	}

	return result
}
