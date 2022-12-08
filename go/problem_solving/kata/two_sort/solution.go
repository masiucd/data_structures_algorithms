package two_sort

import (
	"sort"
	"strings"
)

func twoSort(arr []string) string {
	sort.Strings(arr)
	first := arr[0]
	var res string
	for i, v := range first {
		if i != len(first)-1 {
			res += string(v) + strings.Repeat("*", 3)
		} else {
			res += string(v)
		}
	}
	return res
}

func twoSort2(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}
