package countvowels

var vowels = []rune{'a', 'e', 'i', 'o', 'u'}

func getVowelsCount(input string) (count int) {
	for i := range input {
		for j := range vowels {
			if rune(input[i]) == vowels[j] {
				count++
			}
		}
	}
	return count
}

func getVowelsCount2(input string) (count int) {
	for _, c := range input {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
