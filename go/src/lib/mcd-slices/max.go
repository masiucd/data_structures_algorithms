package mcdslices

import "golang.org/x/exp/constraints"

// Max function takes in slice of integers and return the largest value
func Max[T constraints.Integer](xs []T) T {
	var result T = xs[0]
	for _, v := range xs {
		if v > result {
			result = v
		}
	}
	return result
}
