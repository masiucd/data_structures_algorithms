package group_anagrams

import "sort"

func groupAnagrams(list []string) [][]string {
	var result [][]string
	store := make(map[string][]string)
	for _, v := range list {
		sorted := sortString(v)
		if _, ok := store[sorted]; !ok {
			store[sorted] = []string{v}
		} else {
			store[sorted] = append(store[sorted], v)
		}
	}
	for _, v := range store {
		result = append(result, v)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
	return string(r)
}
