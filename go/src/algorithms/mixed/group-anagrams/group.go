package groupanagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	store := make(map[string][]string)
	for i, word := range strs {
		sortedWord := sortString(word)
		store[sortedWord] = append(store[sortedWord], strs[i])
	}
	for _, list := range store {
		result = append(result, list)
	}
	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
