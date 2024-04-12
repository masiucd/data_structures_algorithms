package checkperfectnumber

func checkPerfectNumber(num int) bool {
	var sum int
	for i := 1; i < num; i++ {
		if num%i == 0 {
			sum += i
		} else if num < sum {
			return false
		}
	}
	return sum == num
}
