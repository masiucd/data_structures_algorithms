package mcdslices

import "golang.org/x/exp/constraints"

// Min function takes in slice of integers and return the smallest value
func Min[T constraints.Integer](xs []T) T {
	var result T = xs[0]
	for _, v := range xs {
		if v < result {
			result = v
		}
	}
	return result
}
