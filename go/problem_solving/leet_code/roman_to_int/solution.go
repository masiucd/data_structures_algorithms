package romantoint

var roman map[string]int = map[string]int{
	"I":  1,
	"V":  5,
	"X":  10,
	"L":  50,
	"C":  100,
	"D":  500,
	"M":  1000,
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanToInt(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[string(s[i:i+2])] > 0 {
			res += roman[string(s[i:i+2])]
			// go to next letter
			i++
		} else {
			res += roman[string(s[i])]
		}
	}
	return res
}
