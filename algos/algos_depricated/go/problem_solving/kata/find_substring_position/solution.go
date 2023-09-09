package find_substring_position

func solution(input, sub string) int {
	var index int
	for index < len(input) {
		if input[index:index+len(sub)] == sub {
			return index
		}
		index += 1
	}
	return -1
}
