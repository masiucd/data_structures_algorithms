package romanencoder

var values []int = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
var romanLetters []string = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

func solution(number int) string {
	var roman string
	for i, n := range values {
		for number >= n {
			number -= n
			roman += romanLetters[i]
		}
	}
	return roman
}
