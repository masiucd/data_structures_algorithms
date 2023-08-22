package at_the_crossroads

func extraNumber(a int, b int, c int) int {
	nums := []int{a, b, c}
	store := make(map[int]int)
	for _, num := range nums {
		if _, ok := store[num]; !ok {
			store[num] = 1
		} else {
			store[num] = store[num] + 1
		}
	}
	for k, v := range store {
		if v == 1 {
			return k
		}
	}

	return -1
}

func extraNumberTwo(a int, b int, c int) int {
	return a ^ b ^ c
}
