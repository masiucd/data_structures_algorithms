package splitstrings

// https://www.codewars.com/kata/515de9ae9dcfc28eb6000001/train/go
func solution(str string) []string {
	size := len(str)
	var res []string
	for i := 0; i < size-1; i += 2 {
		res = append(res, str[i:i+2])
	}
	if size%2 == 0 {
		return res
	}
	res = append(res, string(str[size-1])+"_")
	return res
}
