package commoncharcount

func commonCharacterCount(s1, s2 string) int {
	store1 := make(map[rune]int)
	store2 := make(map[rune]int)

	for _, v := range s1 {
		store1[v]++
	}
	for _, v := range s2 {
		store2[v]++
	}

	var res int
	for k := range store1 {
		if count, ok := store2[k]; ok {

			if count < store1[k] {
				res += count
			} else {
				res += store1[k]
			}
		}
	}

	return res
}
