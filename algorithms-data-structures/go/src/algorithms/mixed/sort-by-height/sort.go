package sortbyheight

import "sort"

func sortByHeight(heights []int) []int {
	var res []int
	var sorted []int
	for _, v := range heights {
		if v != -1 {
			sorted = append(sorted, v)
		}
	}
	sort.Ints(sorted)
	i := -1
	for _, n := range heights {
		if n == -1 {
			res = append(res, -1)
		} else {
			i++
			val := sorted[i]
			res = append(res, val)
		}
	}
	return res
}
